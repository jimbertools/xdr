package service_test

import (
	"testing"

	"github.com/vantorrewannes/watcheye/pkg/watcheye/yara/rules/service"
)

func TestStringRuleService(t *testing.T) {
	stringRulesService, err := service.NewStringRuleFactory()
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
