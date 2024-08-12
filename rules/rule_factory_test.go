package utils

import (
	"testing"
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
	factory := NewStringRuleFactory([]string{abcRule})
	_, err := factory.GetAllRules()
	if err != nil {
		t.Fatalf(`GetAllRules() error = %v`, err)
	}
}

func TestFileRuleFactory(t *testing.T) {
	rulesFilePath := "../test_files/rules.yar"
	factory := NewFileRuleFactory([]string{rulesFilePath})
	_, err := factory.GetAllRules()
	if err != nil {
		t.Fatalf(`GetAllRules() error = %v`, err)
	}
}
