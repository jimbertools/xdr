package journal

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/jimbertools/volmgmt/usn"
	"github.com/jimbertools/volmgmt/usnfilter"
	"github.com/jimbertools/volmgmt/volume"
)

type Journal struct {
	volume  *volume.Volume
	journal *usn.Journal
}

type DiskLetter rune

const (
	A DiskLetter = iota + 65
	B
	C
	D
	E
	F
	G
	H
	I
	J
	K
	L
	M
	N
	O
	P
	Q
	R
	S
	T
	U
	V
	W
	X
	Y
	Z
)

func NewJournal(diskLetter DiskLetter) (*Journal, error) {
	diskLetterString := string(diskLetter)
	volumePath := fmt.Sprintf(`\\.\%s:`, strings.ToUpper(diskLetterString))
	volume, err := volume.New(volumePath)
	if err != nil {
		return nil, err
	}
	journal := Journal{
		volume:  volume,
		journal: volume.Journal(),
	}
	return &journal, nil
}

func (journal *Journal) Close() error {
	return journal.volume.Close()
}

func (journal *Journal) Listen(reason usn.Reason, ctx context.Context, filter usn.Filter, interval time.Duration) (<-chan usn.Record, <-chan error, error)  {
	monitor := journal.journal.Monitor()
	data, err := journal.journal.Query()
	if err != nil {
		return nil, nil, err
	}
	cache, err := journal.journal.Cache(ctx, usnfilter.IsDir, usn.Min, usn.Max)
	if err != nil {
		return nil, nil, err
	}
	usnfilter := func(record usn.Record) bool {
		return (reason == usn.ReasonAny || record.Reason == reason) && (filter == nil || filter(record))
	}
	filer := cache.Filer
	records := monitor.Listen(64)
	errChannel := monitor.Run(data.NextUSN, interval, reason, nil, usnfilter, filer)
	return records, errChannel, nil
}

