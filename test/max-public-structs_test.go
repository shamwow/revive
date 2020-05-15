package test

import (
	"testing"

	"github.com/shamwow/revive/lint"
	"github.com/shamwow/revive/rule"
)

func TestMaxPublicStructs(t *testing.T) {
	testRule(t, "max-public-structs", &rule.MaxPublicStructsRule{}, &lint.RuleConfig{
		Arguments: []interface{}{int64(1)},
	})
}
