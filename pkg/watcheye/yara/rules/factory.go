package rules

import (
	"os"

	"github.com/hillu/go-yara/v4"
)

type YaraRuleFactory struct {
	compiler *yara.Compiler
}

func NewRuleFactory() (*YaraRuleFactory, error) {
	compiler, err := yara.NewCompiler()
	if err != nil {
		return nil, err
	}
	return &YaraRuleFactory{compiler: compiler}, nil
}

func (factory *YaraRuleFactory) AddRuleString(rule string) error {
	return factory.compiler.AddString(rule, "custom")
}

func (factory *YaraRuleFactory) AddRuleFile(filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	return factory.compiler.AddFile(file, "custom")
}

func (factory *YaraRuleFactory) GetRules() (*yara.Rules, error) {
	return factory.compiler.GetRules()
}