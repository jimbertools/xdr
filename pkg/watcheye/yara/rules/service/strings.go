package service

import (
	"github.com/hillu/go-yara/v4"
)

type StringRuleFactory struct {
	compiler *yara.Compiler
}

func NewStringRuleFactory() (*StringRuleFactory, error) {
	compiler, err := yara.NewCompiler()
	if err != nil {
		return nil, err
	}
	return StringRuleFactoryFromCompiler(compiler), nil
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
