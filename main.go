package main

import (
	"context"
	"log"

	"github.com/jimbertools/volmgmt/usn"
	"github.com/vantorrewannes/watcheye/pkg/watcheye/disk/journal"
)

func main() {
	diskLetter := journal.C
	journal, err := journal.NewJournal(diskLetter)
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()
	recordChannel, _, err := journal.Listen(usn.ReasonAny, ctx, nil, 0)
	if err != nil {
		log.Fatal(err)
	}

	for record := range recordChannel {
		log.Println(record.FileName)
	}
}
