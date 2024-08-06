package scanner_test

import (
	"testing"

	"github.com/vantorrewannes/file-scanner/scanner"
	"github.com/vantorrewannes/file-scanner/utils"
)

func TestGetAllRules(t *testing.T) {
	factory := utils.NewRuleFactory()
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
