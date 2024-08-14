package rules_test

import (
	"os"
	"testing"

	"github.com/vantorrewannes/watcheye/pkg/watcheye/yara/rules"
)

func TestFileRuleService(t *testing.T) {
	stringRulesService, err := rules.NewFileRuleFactory()
	if err != nil {
		t.Fatal(err)
	}
	const rulePath = "..\\..\\..\\..\\test\\testdata\\watcheye\\yara\\rules\\abc_rule.yar"
	ruleFile, err := os.Open(rulePath)
	if err != nil {
		t.Fatal(err)
	}
	err = stringRulesService.AddRule(ruleFile)
	if err != nil {
		t.Fatal(err)
	}
	_, err = stringRulesService.Rules()
	if err != nil {
		t.Fatal(err)
	}
}

