package rules_test

import (
	"testing"

	"github.com/vantorrewannes/watcheye/pkg/watcheye/yara/rules"
)

func TestStringRuleFactory(t *testing.T) {
	stringRulesService, err := rules.NewStringRuleFactory()
	if err != nil {
		t.Fatal(err)
	}
	const abcRule = `
	rule abcRule {
			meta: 
				author = "Wannes Vantorre"
			strings:
				$str = "abc"
			condition:
				$str
		}`
	err = stringRulesService.AddRule(abcRule)
	if err != nil {
		t.Fatal(err)
	}
	_, err = stringRulesService.Rules()
	if err != nil {
		t.Fatal(err)
	}
}
