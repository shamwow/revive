package test

import (
	"testing"

	"github.com/shamwow/revive/rule"
)

func TestConfusingResults(t *testing.T) {
	testRule(t, "confusing-results", &rule.ConfusingResultsRule{})
}
