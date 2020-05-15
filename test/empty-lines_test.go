package test

import (
	"testing"

	"github.com/shamwow/revive/rule"
)

// TestEmptyLines rule.
func TestEmptyLines(t *testing.T) {
	testRule(t, "empty-lines", &rule.EmptyLinesRule{})
}
