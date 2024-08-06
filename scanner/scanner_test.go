package scanner_test

import (
	"fmt"
	"testing"

	"github.com/vantorrewannes/file-scanner/scanner"
	"github.com/vantorrewannes/file-scanner/utils"
)

func TestByteScanner(t *testing.T) {
	const abcRule string = `
	rule abcRule {
		meta: 
			author = "Wannes Vantorre"
		strings:
			$str = "abc"
		condition:
			$str
	}`
	factory := utils.NewStringRuleFactory([]string{abcRule})
	rules, err := factory.GetAllRules()
	if err != nil {
		t.Fatalf(`GetAllRules() error = %v`, err)
	}
	scanner := scanner.NewBytesScanner([]byte("abc"))
	_, err = scanner.Scan(rules)
	if err != nil {
		t.Fatalf(`Scan() error = %v`, err)
	}
}

func TestFileScanner(t *testing.T) {
	const abcRule string = `
	rule abcRule {
		meta: 
			author = "Wannes Vantorre"
		strings:
			$str = "abc"
		condition:
			$str
	}`

	const xyzRule string = `
	rule xyzRule {
		meta: 
			author = "Wannes Vantorre"
		strings:
			$str = "xyz"
		condition:
			$str
	}`
	factory := utils.NewStringRuleFactory([]string{abcRule, xyzRule})
	rules, err := factory.GetAllRules()
	if err != nil {
		t.Fatalf(`GetAllRules() error = %v`, err)
	}
	scanner := scanner.NewFileScanner("../test_files/test.txt")
	result, err := scanner.Scan(rules)
	if err != nil {
		t.Fatalf(`Scan() error = %v`, err)
	}
	if len(result) != 2 {
		t.Fatalf(`Scan() error = %v`, fmt.Errorf("expected 2 results, got %d", len(result)))
	}
	if result[0].Rule != "abcRule" {
		t.Fatalf(`Scan() error = %v`, fmt.Errorf("expected abcRule, got %s", result[0].Rule))
	}
	if result[1].Rule != "xyzRule" {
		t.Fatalf(`Scan() error = %v`, fmt.Errorf("expected xyzRule, got %s", result[1].Rule))
	}
}

func TestFilesScanner(t *testing.T) {
	rulesFilePath := "../test_files/rules.yar"
	factory := utils.NewFileRuleFactory([]string{rulesFilePath})
	rules, err := factory.GetAllRules()
	if err != nil {
		t.Fatalf(`GetAllRules() error = %v`, err)
	}
	scanner := scanner.NewFilesScanner([]string{"../test_files/test.txt"})
	result, err := scanner.Scan(rules)
	if err != nil {
		t.Fatalf(`Scan() error = %v`, err)
	}
	if len(result) != 2 {
		t.Fatalf(`Scan() error = %v`, fmt.Errorf("expected 2 results, got %d", len(result)))
	}
	if result[0].Rule != "abcRule" {
		t.Fatalf(`Scan() error = %v`, fmt.Errorf("expected abcRule, got %s", result[0].Rule))
	}
	if result[1].Rule != "xyzRule" {
		t.Fatalf(`Scan() error = %v`, fmt.Errorf("expected xyzRule, got %s", result[1].Rule))
	}
}

func TestDirScanner(t *testing.T) {
	scanner := scanner.NewDirScanner("../test_files/test_dir")
	rulesFilePath := "../test_files/rules.yar"
	factory := utils.NewFileRuleFactory([]string{rulesFilePath})
	rules, err := factory.GetAllRules()
	if err != nil {
		t.Fatalf(`GetAllRules() error = %v`, err)
	}
	result, err := scanner.Scan(rules)
	if err != nil {
		t.Fatalf(`Scan() error = %v`, err)
	}
	if len(result) != 2 {
		t.Fatalf(`Scan() error = %v`, fmt.Errorf("expected 2 results, got %d", len(result)))
	}
	if result[1].Rule != "abcRule" {
		t.Fatalf(`Scan() error = %v`, fmt.Errorf("expected abcRule, got %s", result[0].Rule))
	}
	if result[0].Rule != "xyzRule" {
		t.Fatalf(`Scan() error = %v`, fmt.Errorf("expected xyzRule, got %s", result[1].Rule))
	}
}
