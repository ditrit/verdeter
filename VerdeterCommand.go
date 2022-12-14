package verdeter

import (
	"github.com/ditrit/verdeter/models"
	"github.com/spf13/cobra"
)

// VerdeterCommand is a wrapper around [github.com/spf13/cobra.Command] and viper.
// It provides additional features such as:
//   - support for custom validators to make sure the input data is correct
//   - support for custom constraint functions (see [models.ConstraintFunction])
//   - support for custom normalization function (see [models.NormalizationFunction])
//
// VerdeterCommand make the integration between cobra and viper possible.
type VerdeterCommand struct {
	cmd *cobra.Command

	// parent (if exists)
	parentCmd *VerdeterCommand

	// sub commands
	subCmds map[string]*VerdeterCommand

	// function to call for the command
	runE func(cfg *VerdeterCommand, args []string) error

	// local config keys
	localConfigKeys map[string]*ConfigKey

	// global config keys
	globalConfigKeys map[string]*ConfigKey

	// Required Number of args (no argument allowed by default)
	nbArgs int

	// constraints are predicate functions to be satisfied as a prequisite to command execution
	constraints map[string]models.ConstraintFunction
}

// NewVerdeterCommand is the constructor for VerdeterCommand
// The args "use", "shortDesc" and "longDesc" are string, their role is the same as in [github.com/spf13/cobra.Command]
// The arg runE the callback for cobra
func NewVerdeterCommand(use, shortDesc, longDesc string, runE func(verdeterCmd *VerdeterCommand, args []string) error) *VerdeterCommand {
	var cobraCmd = new(cobra.Command)
	var verdeterCmd = new(VerdeterCommand)

	cobraCmd.PreRunE = preRunCheckE(verdeterCmd)
	cobraCmd.RunE = RunE(verdeterCmd, runE)
	cobraCmd.Use = use
	cobraCmd.Short = shortDesc
	cobraCmd.Long = longDesc

	verdeterCmd.cmd = cobraCmd
	verdeterCmd.runE = runE
	verdeterCmd.subCmds = make(map[string]*VerdeterCommand)
	verdeterCmd.globalConfigKeys = make(map[string]*ConfigKey)
	verdeterCmd.localConfigKeys = make(map[string]*ConfigKey)
	verdeterCmd.constraints = make(map[string]models.ConstraintFunction)

	return verdeterCmd
}

// Add a sub command
func (verdeterCmd *VerdeterCommand) AddSubCommand(sub *VerdeterCommand) {
	verdeterCmd.cmd.AddCommand(sub.cmd)
	verdeterCmd.subCmds[sub.cmd.Name()] = sub
	sub.parentCmd = verdeterCmd
}

// Execute the VerdeterCommand
//
// (panics if called on a subcommand)
func (verdeterCmd *VerdeterCommand) Execute() {
	if verdeterCmd.parentCmd != nil {
		panic("Execute can only be called on root command")
	}
	if err := verdeterCmd.cmd.Execute(); err != nil {
		panic(err)
	}
}

// Lookup returns the ConfigKey structure of the named config key, returning nil if none exists.
func (verdeterCmd *VerdeterCommand) Lookup(configKeyName string) *ConfigKey {
	configKey, ok := verdeterCmd.localConfigKeys[configKeyName]
	if ok {
		return configKey
	}
	configKey, ok = verdeterCmd.globalConfigKeys[configKeyName]
	if ok {
		return configKey
	}
	return nil
}
