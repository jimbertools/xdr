package scanner

import "github.com/hillu/go-yara/v4"

type Scanner interface {
	Scan(rules *yara.Rules) (string, error)
}
