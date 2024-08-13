package main

import (
	"context"
	"log"
	"os"

	disk_watcher "github.com/vantorrewannes/file-scanner/pkg/disk-watcher"
	"github.com/vantorrewannes/file-scanner/pkg/disk-watcher/yara"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Please provide disk letter")
	} else if len(os.Args) < 3 {
		log.Fatal("Please provide YARA rule path")
	}

	disk := os.Args[1]
	yaraRulesPath := os.Args[2]
	

	factory := yara.NewFileRuleFactory([]string{yaraRulesPath})
	rules, err := factory.GetAllRules()
	if err != nil {
		log.Fatalf(`GetAllRules() error = %v`, err)
	}

	diskWatcher := disk_watcher.NewDiskWatcher(rune(disk[0]), rules)
	ctx := context.Background()

	err = diskWatcher.Watch(ctx)
	if err != nil {
		log.Fatal(err)
	}
}