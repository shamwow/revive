package test

import (
	"testing"

	"github.com/shamwow/revive/lint"
	"github.com/shamwow/revive/rule"
)

func TestRangeValAddress(t *testing.T) {
	testRule(t, "range-val-address", &rule.RangeValAddress{}, &lint.RuleConfig{})
}
