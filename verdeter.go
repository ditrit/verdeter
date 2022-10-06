// Package verdeter provides a config system for distributed programs
package verdeter

import (
	"fmt"

	"github.com/ditrit/verdeter/models"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	_ models.ConfigType = iota
	IsStr
	IsInt
	IsBool
	IsUint
)

// return the name of the root command
func (verdeterCmd *VerdeterCommand) GetAppName() string {
	if verdeterCmd.parentCmd != nil {
		return verdeterCmd.parentCmd.GetAppName()
	}
	return verdeterCmd.cmd.Name()
}

// SetNbArgs : function to fix the number of args
func (verdeterCmd *VerdeterCommand) SetNbArgs(nb int) {
	verdeterCmd.nbArgs = nb
}

// Set the validator for a specific config key
func (verdeterCmd *VerdeterCommand) SetValidator(name string, isValid models.Validator) {
	verdeterCmd.isValid[name] = isValid
}

// SetNormalize : function to normalize the value of a config Key (if set)
func (verdeterCmd *VerdeterCommand) SetNormalize(name string, normalize models.NormalizationFunction) {
	verdeterCmd.normalize[name] = normalize
}

// SetDefault : set default value for a key
func (verdeterCmd *VerdeterCommand) SetDefault(name string, value interface{}) {
	viper.SetDefault(name, value)
}

// SetRequired sets a key as required
func (verdeterCmd *VerdeterCommand) SetRequired(name string) {
	verdeterCmd.isRequired[name] = true
}

// SetConstraint sets a constraint
func (verdeterCmd *VerdeterCommand) SetConstraint(msg string, constraint func() bool) {
	verdeterCmd.constraints[msg] = constraint
}

// SetComputedValue sets a value dynamically as the default for a key
func (verdeterCmd *VerdeterCommand) SetComputedValue(name string, fval models.DefaultValueFunction) {
	verdeterCmd.computedValue[name] = fval
}

var fs = afero.NewOsFs()

// InitConfig init Config management
func InitConfig(verdeterCmd *VerdeterCommand) {
	appname := verdeterCmd.GetAppName()
	viper.SetEnvPrefix(appname)

	var configPath = viper.GetString("config_path")
	exists, err := afero.Exists(fs, configPath)
	if err != nil {
		panic(err)
	}
	if !exists {
		panic(
			fmt.Errorf("path %q do not exists", configPath),
		)
	}
	if isDirectory, _ := afero.IsDir(fs, configPath); isDirectory {
		viper.AddConfigPath(configPath)
		viper.SetConfigName(appname)
	} else {
		viper.SetConfigFile(configPath)

	}
	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
			panic("config file not found")
		} else {
			// Config file was found but another error was produced
			panic(fmt.Errorf("config file was found but another error was produced: %s", err.Error()))
		}
	}

}

// Initialize handle initial configuration
func (verdeterCmd *VerdeterCommand) Initialize() {
	cobra.OnInitialize(func() { InitConfig(verdeterCmd) })
}
