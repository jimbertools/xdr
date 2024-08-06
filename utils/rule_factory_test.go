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
	factory := NewRuleFactory([]string{abcRule})
	_, err := factory.GetAllRules()
	if err != nil {
		t.Fatalf(`GetAllRules() error = %v`, err)
	}
}
