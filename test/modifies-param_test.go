package test

import (
	"testing"

	"github.com/shamwow/revive/rule"
)

// TestModifiesParam rule.
func TestModifiesParam(t *testing.T) {
	testRule(t, "modifies-param", &rule.ModifiesParamRule{})
}
