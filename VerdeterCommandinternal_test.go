package verdeter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// This file is intended to hold the tests that need to be run in the verdeter namespace

func TestNewVerdeterCommandAddSubCommand(t *testing.T) {
	verdeterCmd := NewVerdeterCommand(
		"test",
		"test short description",
		"test long description",
		func(verdeterCmd *VerdeterCommand, args []string) error {
			return nil
		},
	)
	assert.Empty(t, verdeterCmd.subCmds)
	verdeterSubCmd := NewVerdeterCommand(
		"subtest",
		"", "",
		func(verdeterCmd *VerdeterCommand, args []string) error {
			return nil
		},
	)
	verdeterCmd.AddSubCommand(verdeterSubCmd)
	assert.NotEmpty(t, verdeterCmd.subCmds)

}
