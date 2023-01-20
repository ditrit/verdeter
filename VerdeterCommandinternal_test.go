package verdeter

import (
	"errors"
	"testing"

	"github.com/spf13/cobra"
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

func TestVerdeterCommandgetRootCommand(t *testing.T) {
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
	verdeterSubCmd.parentCmd = verdeterCmd
	verdeterSubSubCmd := NewVerdeterCommand(
		"subsubtest",
		"", "",
		func(verdeterCmd *VerdeterCommand, args []string) error {
			return nil
		},
	)
	verdeterSubSubCmd.parentCmd = verdeterSubCmd

	assert.Equal(t, verdeterCmd, verdeterSubSubCmd.getRootCommand())
	assert.Equal(t, verdeterCmd, verdeterSubCmd.getRootCommand())
}

func TestBuildVerdeterCommands(t *testing.T) {
	// just some data to fill the structure
	config := VerdeterConfig{
		Use:        "eee",
		Aliases:    []string{"1", "2"},
		SuggestFor: []string{"4", "3"},
		Short:      "short",
		Long:       "long",
		Example:    "exemple",
		ValidArgs:  []string{"valid arg1", "valid arg2"},
		ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			return []string{"valid arg"}, cobra.ShellCompDirectiveFilterDirs
		},
		Args: func(cmd *cobra.Command, args []string) error {
			return errors.New("err args")
		},
		ArgAliases:             []string{"alias1", "alias2"},
		BashCompletionFunction: "BashCompletionFunction",
		Deprecated:             "Deprecated",
		Annotations:            map[string]string{"super": "annotation"},
		Version:                "Version",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
		},
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
		PreRun: func(cmd *cobra.Command, args []string) {
		},
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
		PostRun: func(cmd *cobra.Command, args []string) {
		},
		PostRunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
		PersistentPostRun: func(cmd *cobra.Command, args []string) {
		},
		PersistentPostRunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}
	verdeterCmd := BuildVerdeterCommand(config)

	assert.Equal(t, config.Use, verdeterCmd.cmd.Use)
	assert.Equal(t, config.Aliases, verdeterCmd.cmd.Aliases)
	assert.Equal(t, config.SuggestFor, verdeterCmd.cmd.SuggestFor)
	assert.Equal(t, config.Short, verdeterCmd.cmd.Short)
	assert.Equal(t, config.Long, verdeterCmd.cmd.Long)
	assert.Equal(t, config.Example, verdeterCmd.cmd.Example)
	assert.Equal(t, config.ValidArgs, verdeterCmd.cmd.ValidArgs)
	assert.NotNil(t, config.ValidArgsFunction, verdeterCmd.cmd.ValidArgsFunction)
	assert.NotNil(t, config.Args, verdeterCmd.cmd.Args)
	assert.Equal(t, config.ArgAliases, verdeterCmd.cmd.ArgAliases)
	assert.Equal(t, config.BashCompletionFunction, verdeterCmd.cmd.BashCompletionFunction)
	assert.Equal(t, config.Deprecated, verdeterCmd.cmd.Deprecated)
	assert.Equal(t, config.Annotations, verdeterCmd.cmd.Annotations)
	assert.Equal(t, config.Version, verdeterCmd.cmd.Version)

	assert.NotNil(t, verdeterCmd.cmd.PersistentPreRun)
	assert.NotNil(t, verdeterCmd.cmd.PersistentPreRunE)
	assert.NotNil(t, verdeterCmd.cmd.PreRun)
	assert.NotNil(t, verdeterCmd.cmd.PreRunE)
	assert.NotNil(t, verdeterCmd.cmd.Run)
	assert.NotNil(t, verdeterCmd.cmd.RunE)
	assert.NotNil(t, verdeterCmd.cmd.PostRun)
	assert.NotNil(t, verdeterCmd.cmd.PostRunE)
	assert.NotNil(t, verdeterCmd.cmd.PersistentPostRun)
	assert.NotNil(t, verdeterCmd.cmd.PersistentPostRunE)

	assert.NotNil(t, verdeterCmd.subCmds)
	assert.NotNil(t, verdeterCmd.globalConfigKeys)
	assert.NotNil(t, verdeterCmd.localConfigKeys)
	assert.NotNil(t, verdeterCmd.constraints)
}

func TestBuildVerdeterCommandsNoPreRunE(t *testing.T) {
	// just some data to fill the structure
	config := VerdeterConfig{
		PreRun: func(cmd *cobra.Command, args []string) {
		},
	}
	verdeterCmd := BuildVerdeterCommand(config)
	assert.NotNil(t, verdeterCmd.cmd.PreRun)
	assert.Nil(t, verdeterCmd.cmd.PreRunE)
}
