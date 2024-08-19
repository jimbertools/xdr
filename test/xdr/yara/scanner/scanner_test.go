package scanner_test

import (
	"testing"

	"github.com/jimbertools/xdr/pkg/xdr/yara/scanner"
)

func TestNewYaraScanner(t *testing.T) {
	const ruleFilePath = "..\\..\\..\\..\\test\\testdata\\xdr\\yara\\rules\\abc_rule.yara"
	_, err := scanner.YaraScannerFromRuleFile(ruleFilePath)
	if err != nil {
		t.Fatal(err)
	}
}

func TestYaraScannerScanString(t *testing.T) {
	const ruleFilePath = "..\\..\\..\\..\\test\\testdata\\xdr\\yara\\rules\\abc_rule.yara"
	yaraScanner, err := scanner.YaraScannerFromRuleFile(ruleFilePath)
	if err != nil {
		t.Fatal(err)
	}
	const testString = `This is a test string checking to see if "abc" and "xyz" trigger the yara patterns.`
	matches, err := yaraScanner.ScanString(testString)
	if err != nil {
		t.Fatal(err)
	}
	if len(matches) != 2 {
		t.Fatalf("expected 2 matches, got %d", len(matches))
	}
}

func TestYaraScannerScanFile(t *testing.T) {
	const ruleFilePath = "..\\..\\..\\..\\test\\testdata\\xdr\\yara\\rules\\abc_rule.yara"
	yaraScanner, err := scanner.YaraScannerFromRuleFile(ruleFilePath)
	if err != nil {
		t.Fatal(err)
	}
	const testFilePath = "..\\..\\..\\..\\test\\testdata\\xdr\\yara\\rules\\abc_rule.yara"
	matches, err := yaraScanner.ScanFile(testFilePath)
	if err != nil {
		t.Fatal(err)
	}
	if len(matches) != 2 {
		t.Fatalf("expected 2 matches, got %d", len(matches))
	}
}

func TestYaraScannerScanDir(t *testing.T) {
	const ruleFileDir = "..\\..\\..\\..\\test\\testdata\\xdr\\yara\\rules"
	yaraScanner, err := scanner.YaraScannerFromRuleDir(ruleFileDir)
	if err != nil {
		t.Fatal(err)
	}
	const testFilePath = "..\\..\\..\\..\\test\\testdata\\xdr\\yara\\rules\\abc_rule.yara"
	matches, err := yaraScanner.ScanFile(testFilePath)
	if err != nil {
		t.Fatal(err)
	}
	if len(matches) != 2 {
		t.Fatalf("expected 2 matches, got %d", len(matches))
	}
}
