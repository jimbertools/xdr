package scanner

import (
	"os"

	"github.com/hillu/go-yara/v4"
	"github.com/vantorrewannes/watcheye/pkg/watcheye/yara/rules"
)

type YaraScanner struct {
	rules *yara.Rules
}

func NewYaraScanner(rules *yara.Rules) *YaraScanner {
	return &YaraScanner{rules: rules}
}

func YaraScannerFromRulesFactory(factory rules.RuleFactory) (*YaraScanner, error) {
	yaraRules, err := factory.Rules()
	if err != nil {
		return nil, err
	}
	return NewYaraScanner(yaraRules), nil
}

func YaraScannerFromRuleFile(ruleFilePath string) (*YaraScanner, error) {
	ruleFile, err := os.Open(ruleFilePath)
	if err != nil {
		return nil, err
	}
	fileRulesFactory, err := rules.NewFileRuleFactory()
	if err != nil {
		return nil, err
	}
	err = fileRulesFactory.AddRule(ruleFile)
	if err != nil {
		return nil, err
	}
	return YaraScannerFromRulesFactory(fileRulesFactory)
}

func YaraScannerFromRuleString(ruleString string) (*YaraScanner, error) {
	stringRulesFactory, err := rules.NewStringRuleFactory()
	if err != nil {
		return nil, err
	}
	err = stringRulesFactory.AddRule(ruleString)
	if err != nil {
		return nil, err
	}
	return YaraScannerFromRulesFactory(stringRulesFactory)
}

func (scanner *YaraScanner) ScanFile(filePath string) (yara.MatchRules, error) {
	var matches yara.MatchRules
	err := scanner.rules.ScanFile(filePath, 0, 0, &matches)
	return matches, err
}

func (scanner *YaraScanner) ScanBytes(b []byte) (yara.MatchRules, error) {
	var matches yara.MatchRules
	err := scanner.rules.ScanMem(b, 0, 0, &matches)
	return matches, err
}

func (scanner *YaraScanner) ScanString(str string) (yara.MatchRules, error) { 
	return scanner.ScanBytes([]byte(str))
}
