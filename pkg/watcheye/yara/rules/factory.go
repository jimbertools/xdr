package rules

import (
	"os"

	"github.com/hillu/go-yara/v4"
)

type yaraRuleFactory struct {
	compiler *yara.Compiler
}

func NewRuleFactory() (*yaraRuleFactory, error) {
	compiler, err := yara.NewCompiler()
	if err != nil {
		return nil, err
	}
	return &yaraRuleFactory{compiler: compiler}, nil
}

func (factory *yaraRuleFactory) AddRuleString(rule string) error {
	return factory.compiler.AddString(rule, "custom")
}

func (factory *yaraRuleFactory) AddRuleFile(filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	return factory.compiler.AddFile(file, "custom")
}

func (factory *yaraRuleFactory) GetRules() (*yara.Rules, error) {
	return factory.compiler.GetRules()
}