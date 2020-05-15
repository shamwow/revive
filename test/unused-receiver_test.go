package test

import (
	"testing"

	"github.com/shamwow/revive/rule"
)

func TestUnusedReceiver(t *testing.T) {
	testRule(t, "unused-receiver", &rule.UnusedReceiverRule{})
}
