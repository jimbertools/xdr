package test_yara

import (
	"testing"
	"github.com/vantorrewannes/file-scanner/pkg/file_scanner/yara"
)

func TestStringRuleFactory(t *testing.T) {
	const abcRule string = `
	rule abcRule {
		meta: 
			author = "Wannes Vantorre"
		strings:
			$str = "abc"
		condition:
			$str
	}`
	factory := yara.NewStringRuleFactory([]string{abcRule})
	_, err := factory.GetAllRules()
	if err != nil {
		t.Fatalf(`GetAllRules() error = %v`, err)
	}
}

func TestFileRuleFactory(t *testing.T) {
	rulesFilePath := "../test_files/rules.yar"
	factory := yara.NewFileRuleFactory([]string{rulesFilePath})
	_, err := factory.GetAllRules()
	if err != nil {
		t.Fatalf(`GetAllRules() error = %v`, err)
	}
}
