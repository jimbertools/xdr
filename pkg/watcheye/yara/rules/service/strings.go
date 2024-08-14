package service

import (
	"log"

	"github.com/hillu/go-yara/v4"
)

type StringRuleFactory struct {
	compiler *yara.Compiler
}

func NewStringRuleFactory() *StringRuleFactory {
	compiler, err := yara.NewCompiler()
	if err != nil {
		log.Fatalln(err)
	}
	return StringRuleFactoryFromCompiler(compiler)
}

func StringRuleFactoryFromCompiler(compiler *yara.Compiler) *StringRuleFactory {
	return &StringRuleFactory{compiler: compiler}
}

func (factory *StringRuleFactory) AddRule(rule string) error {
	return factory.compiler.AddString(rule, "custom")
}

func (factory *StringRuleFactory) Rules() (*yara.Rules, error) {
	return factory.compiler.GetRules()
}
