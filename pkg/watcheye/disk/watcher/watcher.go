package watcher

import (
	"context"
	"fmt"
	"time"

	"github.com/hillu/go-yara/v4"
	"github.com/jimbertools/volmgmt/usn"
	"github.com/vantorrewannes/watcheye/pkg/watcheye/disk/journal"
	"github.com/vantorrewannes/watcheye/pkg/watcheye/yara/scanner"
)

type OnYaraMatch func(record usn.Record, path string, matches *yara.MatchRules)
type OnNoYaraMatch func(record usn.Record, path string)

type Watcher struct {
	yaraScanner   *scanner.YaraScanner
	onYaraMatch   OnYaraMatch
	onNoYaraMatch OnNoYaraMatch
}

func NewWatcher(yaraScanner *scanner.YaraScanner, onYaraMatch OnYaraMatch, onNoYaraMatch OnNoYaraMatch) *Watcher {
	return &Watcher{yaraScanner: yaraScanner, onYaraMatch: onYaraMatch, onNoYaraMatch: onNoYaraMatch}
}

func (watcher *Watcher) Watch(diskLetter journal.DiskLetter, reason usn.Reason, ctx context.Context) error {
	journal, err := journal.NewJournal(diskLetter)
	if err != nil {
		return err
	}

	records, errChannel, err := journal.Listen(reason, ctx, nil, time.Millisecond*100)
	if err != nil {
		return err
	}

	const maxTries = 10

	for {
		select {
		case record, ok := <-records:
			if !ok {
				return nil
			}
			go watcher.processRecord(record, diskLetter, maxTries)
		case err, ok := <-errChannel:
			if !ok {
				return err
			}
		}
	}
}

func (watcher *Watcher) processRecord(record usn.Record, diskLetter journal.DiskLetter, maxTries int) {
	recordPath := fmt.Sprintf(`%s:\\%s`, string(diskLetter), record.Path)
	for tries := 0; tries < maxTries; tries++ {
		matches, err := watcher.yaraScanner.ScanFile(recordPath)
		if err == nil {
			if len(matches) > 0 {
				
				if watcher.onYaraMatch != nil {
					watcher.onYaraMatch(record, recordPath, &matches)
				}
			} else {
				if watcher.onNoYaraMatch != nil {
					watcher.onNoYaraMatch(record, recordPath)
				}
			}
			return
		}
		time.Sleep(time.Millisecond * 200)
	}
}
