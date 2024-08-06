package utils

import (
	"log"

	"github.com/hillu/go-yara/v4"
)

type RuleFactory struct {
	compiler *yara.Compiler
}

func NewRuleFactory() *RuleFactory {
	ruleCompiler, err := yara.NewCompiler()
	if ruleCompiler == nil || err != nil {
		log.Fatal("Error to create compiler:", err)
	}

	const abcRule string = `
	rule test {
		meta: 
			author = "Wannes Vantorre"
		strings:
			$str = "abc"
		condition:
			$str
	}`

	const xyzRule string = `
	rule test {
		meta: 
			author = "Wannes Vantorre"
		strings:
			$str = "xyz"
		condition:
			$str
	}`

	rules := []string{abcRule, xyzRule}

	for _, rule := range rules {
		if err = ruleCompiler.AddString(rule, ""); err != nil {
			log.Fatal("Error adding YARA rule:", err)
		}
	}
	return &RuleFactory{ruleCompiler}
}

func (factory *RuleFactory) GetAllRules() (*yara.Rules, error) {
	return factory.compiler.GetRules()
}
