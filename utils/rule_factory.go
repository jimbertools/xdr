package utils

import (
	"log"

	"github.com/hillu/go-yara/v4"
)

type RuleFactory interface {
	*yara.Compiler
	GetAllRules() (*yara.Rules, error)
}

type StringRuleFactory struct {
	compiler *yara.Compiler
}

func NewRuleFactory(rules []string) *StringRuleFactory {
	ruleCompiler, err := yara.NewCompiler()
	if ruleCompiler == nil || err != nil {
		log.Fatal("Error to create compiler:", err)
	}
	for _, rule := range rules {
		if err = ruleCompiler.AddString(rule, ""); err != nil {
			log.Println("Error adding YARA rule:", err)
		}
	}
	return &StringRuleFactory{ruleCompiler}
}

func (factory *StringRuleFactory) GetAllRules() (*yara.Rules, error) {
	return factory.compiler.GetRules()
}
