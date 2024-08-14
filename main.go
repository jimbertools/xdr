package main

import (
	"context"
	"flag"
	"log"
	"path/filepath"

	"github.com/hillu/go-yara/v4"
	"github.com/jimbertools/volmgmt/usn"
	"github.com/jimbertools/xdr/pkg/xdr/disk/journal"
	"github.com/jimbertools/xdr/pkg/xdr/disk/watcher"
	"github.com/jimbertools/xdr/pkg/xdr/yara/scanner"
)

func main() {
	var (
		yaraRulesPath string
		reason        string
	)
	flag.StringVar(&yaraRulesPath, "path", "", "The path to the yara rules file")
	flag.StringVar(&reason, "event", "", "The event that will trigger the watcher to scan the file.")

	flag.Parse()

	if yaraRulesPath == "" {
		log.Fatal("Yara file path is required")
	}
	if reason == "" {
		log.Fatal("Event is required")
	}

	usnReason, err := usn.ParseReason(reason)
	if err != nil {
		log.Fatal(err)
	}
	path, err := filepath.Abs(yaraRulesPath)
	if err != nil {
		log.Fatal(err)
	}
	yaraScanner, err := scanner.YaraScannerFromRuleFile(path)
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()
	watcher := watcher.NewWatcher(yaraScanner, onYaraMatch, onNoYaraMatch)
	err = watcher.Watch(journal.C, usnReason, ctx)
	if err != nil {
		log.Fatal(err)
	}
}

func onYaraMatch(record usn.Record, path string, matches *yara.MatchRules) {
	log.Println("VIRUS:", path)
}

func onNoYaraMatch(record usn.Record, path string) {
	log.Println("CLEAN:", path)
}
