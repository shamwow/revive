package test

import (
	"testing"

	"github.com/shamwow/revive/lint"
	"github.com/shamwow/revive/rule"
)

func TestTypeCasts(t *testing.T) {
	testRule(t, "type-casts", &rule.TypeCastsRule{}, &lint.RuleConfig{
		Arguments: []interface{}{},
	})
}

