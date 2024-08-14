package rules

import "github.com/hillu/go-yara/v4"

type RuleFactory interface {
	Rules() (*yara.Rules, error)
}
