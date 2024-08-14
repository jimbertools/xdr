package rules

import "github.com/hillu/go-yara/v4"

type RuleService interface {
	yara.Compiler
	Rules() (*yara.Rules, error)
}
