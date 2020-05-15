package test

import (
	"testing"

	"github.com/shamwow/revive/lint"
	"github.com/shamwow/revive/rule"
)

func TestFunctionResultsLimit(t *testing.T) {
	testRule(t, "function-result-limit", &rule.FunctionResultsLimitRule{}, &lint.RuleConfig{
		Arguments: []interface{}{int64(3)},
	})
}
