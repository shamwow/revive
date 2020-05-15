package test

import (
	"testing"

	"github.com/shamwow/revive/rule"
)

// String-of-int rule.
func TestStringOfInt(t *testing.T) {
	testRule(t, "string-of-int", &rule.StringOfIntRule{})
}
