package test

import (
	"testing"

	"github.com/shamwow/revive/rule"
)

func TestFlagParam(t *testing.T) {
	testRule(t, "flag-param", &rule.FlagParamRule{})
}
