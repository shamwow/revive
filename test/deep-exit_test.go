package test

import (
	"testing"

	"github.com/shamwow/revive/rule"
)

func TestDeepExit(t *testing.T) {
	testRule(t, "deep-exit", &rule.DeepExitRule{})
	testRule(t, "deep-exit_test", &rule.DeepExitRule{})
}
