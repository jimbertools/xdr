package scanner

import (
	"github.com/hillu/go-yara/v4"
)

type Scanner interface {
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