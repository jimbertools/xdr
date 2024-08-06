package scanner

import (
	"github.com/hillu/go-yara/v4"
)

type Scanner interface {
	Scan(rules *yara.Rules) (yara.MatchRules, error)
}

type BytesScanner struct {
	bytes []byte
}

func (scanner *BytesScanner) Scan(rules *yara.Rules) (yara.MatchRules, error) {
	var matches yara.MatchRules
	err := rules.ScanMem(scanner.bytes, 0, 0, &matches)
	return matches, err
}
