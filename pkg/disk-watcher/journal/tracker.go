package journal

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/jimbertools/volmgmt/usn"
	"github.com/jimbertools/volmgmt/usnfilter"
	Volume "github.com/jimbertools/volmgmt/volume"
)

type DiskTracker struct {
	volume     *Volume.Volume
	journal    *usn.Journal
	process    usn.Processor
	interval   time.Duration
	diskLetter rune
}

// Please close the DiskTracker when done. Only accepts one letter disk names.
func NewDiskTracker(diskLetter rune, processor usn.Processor, interval time.Duration) (DiskTracker, error) {
	volumePath := fmt.Sprintf(`\\.\%s:`, strings.ToUpper(string(diskLetter)))

	volume, err := Volume.New(volumePath)
	if err != nil {
		return DiskTracker{}, err
	}

	diskTracker := DiskTracker{volume, volume.Journal(), processor, interval, diskLetter}
	return diskTracker, nil
}

func (diskTracker *DiskTracker) Track(ctx context.Context, reason usn.Reason) (<-chan usn.Record, <-chan error, error) {
	monitor := diskTracker.journal.Monitor()
	data, err := diskTracker.journal.Query()
	if err != nil {
		return nil, nil, err
	}
	cache, err := diskTracker.journal.Cache(ctx, usnfilter.IsDir, usn.Min, usn.Max)
	if err != nil {
		return nil, nil, err
	}

	filter := func(record usn.Record) bool {
		if reason == usn.ReasonAny {
			return true
		}
		return record.Reason == reason
	}
	filer := cache.Filer
	records := monitor.Listen(64)
	errChannel := monitor.Run(data.NextUSN, diskTracker.interval, reason, nil, filter, filer)
	return records, errChannel, nil
}

func (diskTracker *DiskTracker) Close() {
	diskTracker.journal.Close()
	diskTracker.volume.Close()
}
