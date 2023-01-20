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

	// local config keys
	localConfigKeys map[string]*ConfigKey

	// global config keys
	globalConfigKeys map[string]*ConfigKey

	// constraints are predicate functions to be satisfied as a prequisite to command execution
	constraints map[string]models.ConstraintFunction
}

// VerdeterConfig is meant to be passed to BuildVerdeterCommand() to return an initialized VerdeterCommand
type VerdeterConfig struct {
	// Use is the one-line usage message.
	// Recommended syntax is as follow:
	//   [ ] identifies an optional argument. Arguments that are not enclosed in brackets are required.
	//   ... indicates that you can specify multiple values for the previous argument.
	//   |   indicates mutually exclusive information. You can use the argument to the left of the separator or the
	//       argument to the right of the separator. You cannot use both arguments in a single use of the command.
	//   { } delimits a set of mutually exclusive arguments when one of the arguments is required. If the arguments are
	//       optional, they are enclosed in brackets ([ ]).
	// Example: add [-F file | -D dir]... [-f format] profile
	Use string

	// Aliases is an array of aliases that can be used instead of the first word in Use.
	Aliases []string

	// SuggestFor is an array of command names for which this command will be suggested -
	// similar to aliases but only suggests.
	SuggestFor []string

	// Short is the short description shown in the 'help' output.
	Short string

	// Long is the long message shown in the 'help <this-command>' output.
	Long string

	// Example is examples of how to use the command.
	Example string

	// ValidArgs is list of all valid non-flag arguments that are accepted in shell completions
	ValidArgs []string

	// ValidArgsFunction is an optional function that provides valid non-flag arguments for shell completion.
	// It is a dynamic version of using ValidArgs.
	// Only one of ValidArgs and ValidArgsFunction can be used for a command.
	ValidArgsFunction func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective)

	// Expected arguments
	Args cobra.PositionalArgs

	// ArgAliases is List of aliases for ValidArgs.
	// These are not suggested to the user in the shell completion,
	// but accepted if entered manually.
	ArgAliases []string

	// BashCompletionFunction is custom bash functions used by the legacy bash autocompletion generator.
	// For portability with other shells, it is recommended to instead use ValidArgsFunction
	BashCompletionFunction string

	// Deprecated defines, if this command is deprecated and should print this string when used.
	Deprecated string

	// Annotations are key/value pairs that can be used by applications to identify or
	// group commands.
	Annotations map[string]string

	// Version defines the version for this command. If this value is non-empty and the command does not
	// define a "version" flag, a "version" boolean flag will be added to the command and, if specified,
	// will print content of the "Version" variable. A shorthand "v" flag will also be added if the
	// command does not define one.
	Version string

	// The *Run functions are executed in the following order:
	//   * PersistentPreRun()
	//   * PreRun()
	//   * Run()
	//   * PostRun()
	//   * PersistentPostRun()
	// All functions get the same args, the arguments after the command name.
	//
	// PersistentPreRun: children of this command will inherit and execute.
	PersistentPreRun func(cmd *cobra.Command, args []string)
	// PersistentPreRunE: PersistentPreRun but returns an error.
	PersistentPreRunE func(cmd *cobra.Command, args []string) error
	// PreRun: children of this command will not inherit.
	PreRun func(cmd *cobra.Command, args []string)
	// PreRunE: PreRun but returns an error.
	PreRunE func(cmd *cobra.Command, args []string) error
	// Run: Typically the actual work function. Most commands will only implement this.
	Run func(cmd *cobra.Command, args []string)
	// RunE: Run but returns an error.
	RunE func(cmd *cobra.Command, args []string) error
	// PostRun: run after the Run command.
	PostRun func(cmd *cobra.Command, args []string)
	// PostRunE: PostRun but returns an error.
	PostRunE func(cmd *cobra.Command, args []string) error
	// PersistentPostRun: children of this command will inherit and execute after PostRun.
	PersistentPostRun func(cmd *cobra.Command, args []string)
	// PersistentPostRunE: PersistentPostRun but returns an error.
	PersistentPostRunE func(cmd *cobra.Command, args []string) error
}

// BuildVerdeterCommand takes a VerdeterConfig and return a *VerdeterCommand
func BuildVerdeterCommand(config VerdeterConfig) *VerdeterCommand {
	var cobraCmd = new(cobra.Command)
	var verdeterCmd = new(VerdeterCommand)

	cobraCmd.Use = config.Use
	cobraCmd.Aliases = config.Aliases
	cobraCmd.SuggestFor = config.SuggestFor
	cobraCmd.Short = config.Short
	cobraCmd.Long = config.Long
	cobraCmd.Example = config.Example
	cobraCmd.ValidArgs = config.ValidArgs
	cobraCmd.ValidArgsFunction = config.ValidArgsFunction
	cobraCmd.Args = config.Args
	cobraCmd.ArgAliases = config.ArgAliases
	cobraCmd.BashCompletionFunction = config.BashCompletionFunction
	cobraCmd.Deprecated = config.Deprecated
	cobraCmd.Annotations = config.Annotations
	cobraCmd.Version = config.Version

	cobraCmd.PersistentPreRun = config.PersistentPreRun
	cobraCmd.PersistentPreRunE = config.PersistentPreRunE

	cobraCmd.PreRunE = verdeterCmd.buildPreRunCheckE(config.PreRunE)
	cobraCmd.PreRun = verdeterCmd.buildPreRunCheck(config.PreRun)
	cobraCmd.Run = config.Run
	cobraCmd.RunE = config.RunE
	cobraCmd.PostRun = config.PostRun
	cobraCmd.PostRunE = config.PostRunE
	cobraCmd.PersistentPostRun = config.PersistentPostRun
	cobraCmd.PersistentPostRunE = config.PersistentPostRunE

	verdeterCmd.cmd = cobraCmd
	verdeterCmd.subCmds = make(map[string]*VerdeterCommand)
	verdeterCmd.globalConfigKeys = make(map[string]*ConfigKey)
	verdeterCmd.localConfigKeys = make(map[string]*ConfigKey)
	verdeterCmd.constraints = make(map[string]models.ConstraintFunction)
	return verdeterCmd
}

// NewVerdeterCommand is the constructor for VerdeterCommand
// The args "use", "shortDesc" and "longDesc" are string, their role is the same as in [github.com/spf13/cobra.Command]
// The arg runE the callback for cobra
func NewVerdeterCommand(use, shortDesc, longDesc string, runE func(verdeterCmd *VerdeterCommand, args []string) error) *VerdeterCommand {
	var cobraCmd = new(cobra.Command)
	var verdeterCmd = new(VerdeterCommand)

	cobraCmd.PreRunE = verdeterCmd.buildPreRunCheckE(
		func(cobraCmd *cobra.Command, args []string) error {
			return nil
		})
	cobraCmd.RunE = func(verdeterCommand *VerdeterCommand) func(cmd *cobra.Command, args []string) error {
		return func(cmd *cobra.Command, args []string) error {
			return runE(verdeterCommand, args)
		}
	}(verdeterCmd)
	cobraCmd.Use = use
	cobraCmd.Short = shortDesc
	cobraCmd.Long = longDesc

	verdeterCmd.cmd = cobraCmd
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

// return the root of the command graph
func (verdeterCmd *VerdeterCommand) getRootCommand() *VerdeterCommand {
	if verdeterCmd.parentCmd != nil {
		return verdeterCmd.parentCmd.getRootCommand()
	}
	return verdeterCmd
}
