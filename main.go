package main

import (
	"context"
	"log"

	"github.com/hillu/go-yara/v4"
	"github.com/jimbertools/volmgmt/usn"
	"github.com/vantorrewannes/watcheye/pkg/watcheye/disk/journal"
	"github.com/vantorrewannes/watcheye/pkg/watcheye/disk/watcher"
	"github.com/vantorrewannes/watcheye/pkg/watcheye/yara/scanner"
)

func main() {
	const ruleFilePath = "rules.yar"
	yaraScanner, err := scanner.YaraScannerFromRuleFile(ruleFilePath)
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()
	watcher := watcher.NewWatcher(yaraScanner, onYaraMatch, onNoYaraMatch)
	err = watcher.Watch(journal.C, usn.ReasonFileCreate, ctx)
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
