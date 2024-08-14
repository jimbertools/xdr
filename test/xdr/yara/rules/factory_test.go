package rules_test

import (
	"testing"

	"github.com/jimbertools/xdr/pkg/xdr/yara/rules"
)

func TestNewRuleFactory(t *testing.T) {
	_, err := rules.NewRuleFactory()
	if err != nil {
		t.Fatal(err)
	}
}

func TestAddRule(t *testing.T) {
	yaraRuleFactory, err := rules.NewRuleFactory()
	if err != nil {
		t.Fatal(err)
	}
	const ruleFilePath = "..\\..\\..\\..\\test\\testdata\\xdr\\yara\\rules\\abc_rule.yar"
	err = yaraRuleFactory.AddRuleFile(ruleFilePath)
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetRules(t *testing.T) {
	yaraRuleFactory, err := rules.NewRuleFactory()
	if err != nil {
		t.Fatal(err)
	}
	const ruleFilePath = "..\\..\\..\\..\\test\\testdata\\xdr\\yara\\rules\\abc_rule.yar"
	err = yaraRuleFactory.AddRuleFile(ruleFilePath)
	if err != nil {
		t.Fatal(err)
	}
	_, err = yaraRuleFactory.GetRules()
	if err != nil {
		t.Fatal(err)
	}
}