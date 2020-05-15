package test

import (
	"testing"

	"github.com/shamwow/revive/lint"
	"github.com/shamwow/revive/rule"
)

func TestArgumentLimit(t *testing.T) {
	testRule(t, "argument-limit", &rule.ArgumentsLimitRule{}, &lint.RuleConfig{
		Arguments: []interface{}{int64(3)},
	})
}
