package test

import (
	"testing"

	"github.com/shamwow/revive/rule"
)

// Atomic rule.
func TestAtomic(t *testing.T) {
	testRule(t, "atomic", &rule.AtomicRule{})
}
