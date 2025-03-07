package test

import (
	"testing"

	"github.com/shamwow/revive/lint"
	"github.com/shamwow/revive/rule"
)

func TestVarNaming(t *testing.T) {
	testRule(t, "var-naming", &rule.VarNamingRule{}, &lint.RuleConfig{
		Arguments: []interface{}{[]interface{}{"ID"}, []interface{}{"VM"}},
	})
}
