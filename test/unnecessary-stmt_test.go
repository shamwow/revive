package test

import (
	"testing"

	"github.com/shamwow/revive/rule"
)

// TestUnnecessaryStmt rule.
func TestUnnecessaryStmt(t *testing.T) {
	testRule(t, "unnecessary-stmt", &rule.UnnecessaryStmtRule{})
}
