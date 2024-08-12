package main

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	LibYara "github.com/hillu/go-yara/v4"
	"github.com/jimbertools/volmgmt/usn"
	Yara "github.com/vantorrewannes/file-scanner/pkg/file_scanner/yara"
	"github.com/vantorrewannes/file-scanner/tracker"
)

func main() {
	rulesFilePath := "yara-rules-core.yar"
	factory := Yara.NewFileRuleFactory([]string{rulesFilePath})
	rules, err := factory.GetAllRules()
	if err != nil {
		log.Fatalf(`GetAllRules() error = %v`, err)
	}

	diskTracker, err := tracker.NewDiskTracker("C", filterRecords, nil, time.Second)
	if err != nil {
		log.Fatal(err)
	}
	defer diskTracker.Close()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	filePathsChannel := make(chan string)

	go printPaths(filePathsChannel, rules)

	err = diskTracker.Watch(ctx, usn.ReasonAny, filePathsChannel)
	if err != nil {
		log.Fatal(err)
	}
}

func printPaths(filePathsChannel chan string, rules *LibYara.Rules) {
	fmt.Println("STARTED")
	for filePath := range filePathsChannel {
		fileScanner := Yara.NewFileScanner(filePath)
		matches, err := fileScanner.Scan(rules)
		if err != nil && err.Error() == "could not open file" {
			fmt.Println("LOCKED: ", filePath)
		} else if err != nil {
			fmt.Println(err)
		}
		if err == nil && len(matches) > 0 {
			fmt.Println("HIT: ", filePath)
		} else if err == nil {
			fmt.Println("MISS: ", filePath)
		}
	}
}

func filterRecords(record usn.Record) bool {
	return strings.HasPrefix(record.Path, "Users\\Wannes\\Desktop")
}

// import (
// 	"context"
// 	"fmt"
// 	"os"
// 	"os/signal"

// 	"syscall"
// 	"time"

// 	"github.com/jimbertools/volmgmt/usn"
// 	"github.com/jimbertools/volmgmt/usnfilter"
// 	"golang.org/x/sys/windows"
// )

// func main() {

// 	path := "C:"

// 	reason := usn.ReasonAny

// 	journal, err := usn.NewJournal(path)
// 	if err != nil {
// 		fmt.Printf("Unable to create monitor: %v\n", err)
// 		os.Exit(2)
// 	}
// 	defer journal.Close()

// 	data, err := journal.Query()
// 	if err == windows.ERROR_JOURNAL_NOT_ACTIVE {
// 		fmt.Print("USN Journal is not active. Creating new journal...\n")
// 		err = journal.Create(0, 0)
// 	}

// 	if err != nil {
// 		fmt.Printf("Unable to access USN Journal: %v\n", err)
// 		os.Exit(2)
// 	}

// 	monitor := journal.Monitor()
// 	defer monitor.Close()

// 	feed := monitor.Listen(64) // Register the feed before starting the monitor

// 	cache, err := journal.Cache(context.Background(), usnfilter.IsDir, 0, data.NextUSN)
// 	if err != nil {
// 		fmt.Printf("Journal cache error: %v\n", err)
// 		os.Exit(2)
// 	}

// 	cacheUpdater := func(record usn.Record) {
// 		if usnfilter.IsDir(record) {
// 			cache.Set(record)
// 		}
// 	}
// 	errC := monitor.Run(data.NextUSN, time.Millisecond*100, reason, cacheUpdater, nil, cache.Filer)

// 	done := make(chan struct{})
// 	go run(feed, done)

// 	ch := make(chan os.Signal, 1)
// 	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)

// 	select {
// 	case <-ch:
// 	case <-done:
// 	case err = <-errC:
// 		if err != nil {
// 			fmt.Printf("monitor USN Journal: %v\n", err)
// 		}
// 	}
// }

// func run(feed <-chan usn.Record, done chan struct{}) {
// 	defer close(done)

// 	for record := range feed {
// 		fmt.Println("PATH: ", record.Path)
// 	}
// }

// package main

// import (
// 	"context"
// 	"fmt"

// 	// "fmt"
// 	"log"
// 	"time"

// 	"github.com/jimbertools/volmgmt/usn"
// 	"github.com/vantorrewannes/file-scanner/scanner"
// 	"github.com/vantorrewannes/file-scanner/tracker"
// 	"github.com/vantorrewannes/file-scanner/utils"
// )

// func main() {
// 	tracker, err := tracker.NewDiskTracker("C", nil, nil, time.Hour)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer tracker.Close()

// 	filePathsChannel := make(chan string)

// 	ctx, cancel := context.WithCancel(context.Background())
// 	defer cancel()

// 	rulesFilePath := "yara-rules-core.yar"
// 	factory := utils.NewFileRuleFactory([]string{rulesFilePath})
// 	rules, err := factory.GetAllRules()
// 	if err != nil {
// 		log.Fatalf(`GetAllRules() error = %v`, err)
// 	}

// 	go func() {
// 		err := tracker.Watch(ctx, usn.ReasonAny, filePathsChannel)
// 		if err != nil {
// 			log.Println(err)
// 		}
// 	}()

// 	for filePath := range filePathsChannel {

// 		fileScanner := scanner.NewFileScanner(filePath)
// 		_, err := fileScanner.Scan(rules)
// 		if err != nil {
// 			// log.Println(err)
// 		}
// 		fmt.Println("FOUND: ", filePath)
// 		// if len(matches) > 0 {
// 		// 	fmt.Println("HIT: ", filePath)
// 		// } else {
// 		// 	fmt.Println("MISS: ", filePath)
// 		// }
// 	}
// }

// // import (
// // 	"fmt"

// // 	"github.com/shirou/gopsutil/disk"
// // )

// // func main() {
// // 	partitions, _ := disk.Partitions(true)
// // 	for _, partition := range partitions {
// // 		fmt.Println(partition.Mountpoint)
// // 	}
// // }
