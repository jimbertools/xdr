package service

import (
	"os"
	"github.com/hillu/go-yara/v4"
)

type FileRuleFactory struct {
	compiler *yara.Compiler
}

func NewFileRuleFactory() (*FileRuleFactory, error) {
	compiler, err := yara.NewCompiler()
	if err != nil {
		return nil, err
	}
	return FileRuleFactoryFromCompiler(compiler), nil
}

func FileRuleFactoryFromCompiler(compiler *yara.Compiler) *FileRuleFactory {
	return &FileRuleFactory{compiler: compiler}
}

func (factory *FileRuleFactory) AddRule(file *os.File) error {
	return factory.compiler.AddFile(file, "custom")
}

func (factory *FileRuleFactory) Rules() (*yara.Rules, error) {
	return factory.compiler.GetRules()
}
