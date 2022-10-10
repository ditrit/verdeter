// Package verdeter provides a config system for distributed programs
package verdeter

import (
	"errors"
	"fmt"
	"path"
	"strings"

	"github.com/ditrit/verdeter/models"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/afero"
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

var ErrConfigFileNotFound = errors.New("config file not found")

// initConfig init Config management
func initConfig(verdeterCmd *VerdeterCommand) error {
	appname := verdeterCmd.GetAppName()
	viper.SetEnvPrefix(appname)
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	var locations []string
	var configPath = viper.GetString("config_path")
	if configPath != "" {
		locations = append(locations, configPath)
	}

	locations = append(locations, ".")
	homeFolderLocation, err := homedir.Dir()
	if err == nil {
		locations = append(locations, path.Join(homeFolderLocation, ".config", appname)+"/")
	}
	locations = append(locations,
		fmt.Sprintf("/etc/%s/", appname),
	)
	for _, location := range locations {
		err = tryPath(location, appname)
		if err != nil {
			if errors.Is(err, ErrConfigFileNotFound) {
				continue
			}
			return err
		}
		break

	}
	return nil
}

// try to read the configuration at the config path
func tryPath(configPath string, appname string) error {
	exists, err := afero.Exists(fs, configPath)
	if err != nil {
		return err
	}
	if !exists {
		return ErrConfigFileNotFound
	}
	if isDirectory, _ := afero.IsDir(fs, configPath); isDirectory {
		viper.AddConfigPath(configPath)
		viper.SetConfigName(appname)
	} else {
		viper.SetConfigFile(configPath)
	}

	err = viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return ErrConfigFileNotFound
		}
		return fmt.Errorf("config file was found but another error was produced: %s", err.Error())
	}
	return nil
}
