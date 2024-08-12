package yara

import (
	"os"
	"path/filepath"

	"github.com/hillu/go-yara/v4"
	"golang.org/x/sync/errgroup"
)

type YaraScanner interface {
	Scan(rules *yara.Rules) (yara.MatchRules, error)
}

type BytesScanner struct {
	bytes []byte
}

func NewBytesScanner(bytes []byte) *BytesScanner {
	return &BytesScanner{
		bytes: bytes,
	}
}

func (scanner *BytesScanner) Scan(rules *yara.Rules) (yara.MatchRules, error) {
	var matches yara.MatchRules
	err := rules.ScanMem(scanner.bytes, 0, 0, &matches)
	return matches, err
}

type FileScanner struct {
	filePath string
}

func NewFileScanner(filePath string) *FileScanner {
	return &FileScanner{
		filePath: filePath,
	}
}

func (scanner *FileScanner) Scan(rules *yara.Rules) (yara.MatchRules, error) {
	var matches yara.MatchRules
	err := rules.ScanFile(scanner.filePath, 0, 0, &matches)
	return matches, err
}

type FilesScanner struct {
	fileScanners []FileScanner
}

func NewFilesScanner(filePaths []string) *FilesScanner {
	fileScanners := []FileScanner{}
	for _, filePath := range filePaths {
		fileScanners = append(fileScanners, *NewFileScanner(filePath))
	}
	return &FilesScanner{
		fileScanners: fileScanners,
	}
}

func (scanner *FilesScanner) Scan(rules *yara.Rules) (yara.MatchRules, error) {
	matchesChannel := make(chan yara.MatchRules)
	errorGroup := errgroup.Group{}
	var err error
	for _, fileScanner := range scanner.fileScanners {
		errorGroup.Go(func() error {
			ruleMatches, err := fileScanner.Scan(rules)
			if err != nil {
				return err
			}
			matchesChannel <- ruleMatches
			return nil
		})
	}
	go func() {
		err = errorGroup.Wait()
		close(matchesChannel)
	}()
	var matches yara.MatchRules
	for ruleMatch := range matchesChannel {
		matches = append(matches, ruleMatch...)
	}
	return matches, err
}

type DirScanner struct {
	dirPath string
}

func NewDirScanner(dirPath string) *DirScanner {
	return &DirScanner{
		dirPath: dirPath,
	}
}

func (scanner *DirScanner) Scan(rules *yara.Rules) (yara.MatchRules, error) {
	dirEntries, err := os.ReadDir(scanner.dirPath)
	if err != nil {
		return nil, err
	}
	var matches yara.MatchRules
	var filePaths []string
	for _, dirEntry := range dirEntries {
		path := filepath.Join(scanner.dirPath, dirEntry.Name())
		if dirEntry.IsDir() {
			dirScanner := NewDirScanner(path)
			ruleMatches, err := dirScanner.Scan(rules)
			if err != nil {
				return nil, err
			}
			matches = append(matches, ruleMatches...)
		} else {
			filePaths = append(filePaths, path)
		}
	}
	fileScanner := NewFilesScanner(filePaths)
	ruleMatches, err := fileScanner.Scan(rules)
	if err != nil {
		return nil, err
	}
	matches = append(matches, ruleMatches...)
	return matches, nil
}
