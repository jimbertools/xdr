package scanner

import (
	"github.com/hillu/go-yara/v4"
	"github.com/jimbertools/xdr/pkg/xdr/yara/rules"
)

type YaraScanner struct {
	rules *yara.Rules
}

func NewYaraScanner(rules *yara.Rules) *YaraScanner {
	return &YaraScanner{rules: rules}
}

func YaraScannerFromRulesFactory(factory rules.YaraRuleFactory) (*YaraScanner, error) {
	yaraRules, err := factory.GetRules()
	if err != nil {
		return nil, err
	}
	return NewYaraScanner(yaraRules), nil
}

func YaraScannerFromRuleFile(ruleFilePath string) (*YaraScanner, error) {
	yaraRuleFactory, err := rules.NewRuleFactory()
	if err != nil {
		return nil, err
	}
	err = yaraRuleFactory.AddRuleFile(ruleFilePath)
	if err != nil {
		return nil, err
	}
	return YaraScannerFromRulesFactory(*yaraRuleFactory)
}

func YaraScannerFromRuleString(ruleString string) (*YaraScanner, error) {
	yaraRuleFactory, err := rules.NewRuleFactory()
	if err != nil {
		return nil, err
	}
	err = yaraRuleFactory.AddRuleString(ruleString)
	if err != nil {
		return nil, err
	}
	return YaraScannerFromRulesFactory(*yaraRuleFactory)
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
