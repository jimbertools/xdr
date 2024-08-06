package utils

import (
	"log"
	"os"

	"github.com/hillu/go-yara/v4"
)

type RuleFactory interface {
	*yara.Compiler
	GetAllRules() (*yara.Rules, error)
}

type StringRuleFactory struct {
	compiler *yara.Compiler
}

func NewStringRuleFactory(rules []string) *StringRuleFactory {
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

type FileRuleFactory struct {
	compiler *yara.Compiler
}

func NewFileRuleFactory(rulesFilePaths []string) *FileRuleFactory {
	ruleCompiler, err := yara.NewCompiler()
	if ruleCompiler == nil || err != nil {
		log.Fatal("Error to create compiler:", err)
	}
	for _, ruleFilePath := range rulesFilePaths {
		ruleFile, err := os.Open(ruleFilePath)
		if err != nil {
			log.Println("Error opening YARA rule file:", err)
			continue
		}
		if err = ruleCompiler.AddFile(ruleFile, ""); err != nil {
			log.Println("Error adding YARA rule:", err)
		}
	}
	return &FileRuleFactory{ruleCompiler}
}

func (factory *FileRuleFactory) GetAllRules() (*yara.Rules, error) {
	return factory.compiler.GetRules()
}
