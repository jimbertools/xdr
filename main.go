package main

import (
	"context"
	"log"

	disk_watcher "github.com/vantorrewannes/file-scanner/pkg/disk-watcher"
	"github.com/vantorrewannes/file-scanner/pkg/disk-watcher/yara"
)

func main() {
	rulesFilePath := "test\\testdata\\yara-rules\\rules.yar"
	factory := yara.NewFileRuleFactory([]string{rulesFilePath})
	rules, err := factory.GetAllRules()
	if err != nil {
		log.Fatalf(`GetAllRules() error = %v`, err)
	}

	diskWatcher := disk_watcher.NewDiskWatcher('C', rules)
	ctx := context.Background()

	err = diskWatcher.Watch(ctx)
	if err != nil {
		log.Fatal(err)
	}
}

// package main

// import (
// 	"context"
// 	"fmt"
// 	"log"
// 	"sync"
// 	"time"

// 	"github.com/jimbertools/volmgmt/usn"
// 	"github.com/vantorrewannes/file-scanner/pkg/file_scanner/journal"
// )

// func main() {
// 	diskTracker, err := journal.NewDiskTracker('C', nil, time.Second)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer diskTracker.Close()

// 	recordChannel, errorChannel, err := diskTracker.Track(context.Background(), usn.ReasonFileCreate)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	wg := sync.WaitGroup{}

// 	wg.Add(1)
// 	go printRecords(recordChannel)
// 	go printRecordErrors(errorChannel)
// 	wg.Wait()
// 	fmt.Println("HIT END OF PROGRAM")
// }

// func printRecords(recordChannel <-chan usn.Record) {
// 	for record := range recordChannel {
// 		fmt.Println(record.Path)
// 	}
// }

// func printRecordErrors(errorChannel <-chan error) {
// 	for err := range errorChannel {
// 		fmt.Println(err)
// 	}
// }
