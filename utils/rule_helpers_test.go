package utils

import (
	"testing"
)

func TestGetAllRules(t *testing.T) {
	factory := NewRuleFactory()
	rules := factory.GetAllRules()
	if len(rules) != 2 {
		t.Fatalf(`len(GetAllRules()) = %d, want 2`, len(rules))
	}
}
