package test

import (
	"testing"

	"github.com/shamwow/revive/rule"
)

// IdenticalBranches rule.
func TestIdenticalBranches(t *testing.T) {
	testRule(t, "identical-branches", &rule.IdenticalBranchesRule{})
}
