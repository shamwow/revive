package test

import (
	"testing"

	"github.com/shamwow/revive/lint"
	"github.com/shamwow/revive/rule"
)

func TestRangeValInClosure(t *testing.T) {
	testRule(t, "range-val-in-closure", &rule.RangeValInClosureRule{}, &lint.RuleConfig{})
}
