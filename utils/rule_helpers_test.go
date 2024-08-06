package utils

import (
	"testing"
)

func TestGetAllRules(t *testing.T) {
	factory := NewRuleFactory()
	_, err := factory.GetAllRules()
	if err != nil {
		t.Fatalf(`GetAllRules() error = %v`, err)
	}
}
