package test

import (
	"testing"

	"github.com/shamwow/revive/rule"
)

// BoolLiteral rule.
func TestBoolLiteral(t *testing.T) {
	testRule(t, "bool-literal-in-expr", &rule.BoolLiteralRule{})
}
