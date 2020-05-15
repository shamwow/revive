package test

import (
	"testing"

	"github.com/shamwow/revive/lint"
	"github.com/shamwow/revive/rule"
)

func TestCognitiveComplexity(t *testing.T) {
	testRule(t, "cognitive-complexity", &rule.CognitiveComplexityRule{}, &lint.RuleConfig{
		Arguments: []interface{}{int64(0)},
	})
}
