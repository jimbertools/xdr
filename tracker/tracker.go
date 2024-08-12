package tracker

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/jimbertools/volmgmt/usn"
	Usn "github.com/jimbertools/volmgmt/usn"
	"github.com/jimbertools/volmgmt/usnfilter"
	Volume "github.com/jimbertools/volmgmt/volume"
	"golang.org/x/sync/errgroup"
)

type DiskTracker struct {
	volume   *Volume.Volume
	journal  *Usn.Journal
	filter   Usn.Filter
	process Usn.Processor
	interval time.Duration
}

// Please close the DiskTracker when done. Only accepts one letter disk names.
func NewDiskTracker(diskLetter string, filter Usn.Filter, processor Usn.Processor,  interval time.Duration) (DiskTracker, error) {
	volumePath := fmt.Sprintf(`\\.\%s:`, strings.ToUpper(diskLetter))

	volume, err := Volume.New(volumePath)
	if err != nil {
		return DiskTracker{}, err
	}

	diskTracker := DiskTracker{volume, volume.Journal(), filter, processor, interval}
	return diskTracker, nil
}

func (diskTracker *DiskTracker) Watch(ctx context.Context, reason usn.Reason, filePaths chan<- string) error {
	monitor := diskTracker.journal.Monitor()
	defer monitor.Close()

	cache, err := diskTracker.journal.Cache(ctx, usnfilter.IsDir, usn.Min, usn.Max)
	if err != nil {
		return err
	}

	filer := cache.Filer
	records := monitor.Listen(64)
	errChannel := monitor.Run(usn.Min, diskTracker.interval, reason, nil, diskTracker.filter, filer)

	errorGroup := errgroup.Group{}
	errorGroup.Go(func() error {
		for {
			select {
			case <-ctx.Done():
				return ctx.Err()
			case err := <-errChannel:
				return err
			case record := <-records:
				filePaths <- "C:\\" + record.Path
			}
		}
	})

	return errorGroup.Wait()
}

func (diskTracker *DiskTracker) Monitor() *Usn.Monitor {
	return diskTracker.journal.Monitor()
}

func (diskTracker *DiskTracker) Close() {
	diskTracker.journal.Close()
	diskTracker.volume.Close()
}
