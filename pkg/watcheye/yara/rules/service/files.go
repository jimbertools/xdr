package service

import (
	"log"
	"os"
	"github.com/hillu/go-yara/v4"
)

type FileRuleFactory struct {
	compiler *yara.Compiler
}

func NewFileRuleFactory() *FileRuleFactory {
	compiler, err := yara.NewCompiler()
	if err != nil {
		log.Fatalln(err)
	}
	return FileRuleFactoryFromCompiler(compiler)
}

func FileRuleFactoryFromCompiler(compiler *yara.Compiler) *FileRuleFactory {
	return &FileRuleFactory{compiler: compiler}
}

func (factory *FileRuleFactory) AddRule(path *os.File) error {
	return factory.compiler.AddFile(path, "custom")
}

func (factory *FileRuleFactory) Rules() (*yara.Rules, error) {
	return factory.compiler.GetRules()
}
