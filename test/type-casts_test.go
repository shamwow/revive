package test

import (
	"testing"

	"github.com/mgechev/revive/lint"
	"github.com/mgechev/revive/rule"
)

func TestTypeCasts(t *testing.T) {
	testRule(t, "type-casts", &rule.TypeCastsRule{}, &lint.RuleConfig{
		Arguments: []interface{}{},
	})
}

