// Package verdeter provides a config system for distributed programs
package verdeter

import (
	"fmt"

	"strconv"

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
)

// return the name of the root command
func (vc *VerdeterCommand) GetAppName() string {
	if vc.parentCmd != nil {
		return vc.parentCmd.GetAppName()
	}
	return vc.cmd.Name()
}

// GetInstanceName return the name of the instance of the current app node
// The instance name in composed of the app name joined by the offset is the key is set by the developer or the environment
//
// Example:
//   - if the offset is set
//   - if the offset not set or inferior or equal to 0
func (vc *VerdeterCommand) GetInstanceName() string {
	appName := vc.GetAppName()
	offset := GetOffset()
	if offset != 0 {
		appName = appName + "_" + strconv.Itoa(offset)
	}
	return appName
}

// SetNbArgs : function to fix the number of args
func (vc *VerdeterCommand) SetNbArgs(nb int) {
	vc.nbArgs = nb
}

// Set the validator for a specific config key
func (vc *VerdeterCommand) SetValidator(name string, isValid models.Validator) {
	vc.isValid[name] = isValid
}

// SetNormalize : function to normalize the value of a config Key (if set)
func (vc *VerdeterCommand) SetNormalize(name string, normalize models.NormalizationFunction) {
	vc.normalize[name] = normalize
}

// SetDefault : set default value for a key
func (vc *VerdeterCommand) SetDefault(name string, value interface{}) {
	viper.SetDefault(name, value)
}

// SetRequired sets a key as required
func (vc *VerdeterCommand) SetRequired(name string) {
	vc.isRequired[name] = true
}

// SetConstraint sets a constraint
func (vc *VerdeterCommand) SetConstraint(msg string, constraint func() bool) {
	vc.constraints[msg] = constraint
}

// SetComputedValue sets a value dynamically as the default for a key
func (vc *VerdeterCommand) SetComputedValue(name string, fval models.DefaultValueFunction) {
	vc.computedValue[name] = fval
}

var fs = afero.NewOsFs()

// InitConfig init Config management
func InitConfig(vc *VerdeterCommand) {
	initOffset(vc)
	instanceName := vc.GetInstanceName()
	viper.SetEnvPrefix(instanceName)

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
		viper.SetConfigName(instanceName)
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
func (vc *VerdeterCommand) Initialize() {
	cobra.OnInitialize(func() { InitConfig(vc) })
}
