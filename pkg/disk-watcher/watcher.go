package disk_watcher

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	Yara "github.com/hillu/go-yara/v4"
	"github.com/jimbertools/volmgmt/usn"
	"github.com/vantorrewannes/file-scanner/pkg/disk-watcher/journal"
	"github.com/vantorrewannes/file-scanner/pkg/disk-watcher/yara"
)

type DiskWatcher struct {
	diskLetter rune
	yaraRules  *Yara.Rules
}

func NewDiskWatcher(diskLetter rune, yaraRules *Yara.Rules) *DiskWatcher {
	return &DiskWatcher{
		diskLetter: diskLetter,
		yaraRules:  yaraRules,
	}
}

func (diskWatcher *DiskWatcher) Watch(ctx context.Context) error {
	diskTracker, err := journal.NewDiskTracker('C', nil, time.Second)
	if err != nil {
		log.Fatal(err)
	}
	defer diskTracker.Close()

	recordChannel, errorChannel, err := diskTracker.Track(context.Background(), usn.ReasonAny)
	if err != nil {
		log.Fatal(err)
	}

	wg := sync.WaitGroup{}
	wg.Add(1)

	for record := range recordChannel {
		go diskWatcher.processRecord(record, 10)
	}

	for err := range errorChannel {
		return err
	}
	wg.Wait()
	return nil
}

func (diskWatcher *DiskWatcher) processRecord(record usn.Record, maxTries int) {
	path := fmt.Sprintf(`%s:\\%s`, string(diskWatcher.diskLetter), record.Path)
	for tries := 0; tries < maxTries; tries++ {
		fileScanner := yara.NewFileScanner(path)
		matches, err := fileScanner.Scan(diskWatcher.yaraRules)

		if err == nil {
			if len(matches) > 0 {
				log.Println("VIRUS:", path)
			} else {
				log.Println("CLEAN:", path)
			}
			return
		}
		time.Sleep(time.Millisecond * 200)
	}
}
