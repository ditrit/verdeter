package verdeter

import (
	"fmt"

	"github.com/ditrit/verdeter/models"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

// LKey defines a local flag for cobra bound to env and config file
//
// For an explanation of the difference
func (vc *VerdeterCommand) LKey(name string, valType models.ConfigType, short string, usage string) error {
	vc.keyType[name] = valType
	return Key(vc.cmd, name, valType, short, usage, false)
}

// GKey defines a global flag for cobra bound to env and config file
func (vc *VerdeterCommand) GKey(name string, valType models.ConfigType, short string, usage string) error {
	vc.keyType[name] = valType
	return Key(vc.cmd, name, valType, short, usage, true)
}

// Key defines a flag in cobra bound to env and files
//
// for booleans, the default value will be set to false
// for strings, the default value will be set to ""
// for floats, the default value will be set to 0
func Key(cmd *cobra.Command, name string, valueType models.ConfigType, short string, usage string, global bool) error {
	var flagSet = cmd.PersistentFlags()
	if !global {
		flagSet = cmd.Flags()
	}
	switch valueType {
	case IsStr:
		keyString(flagSet, name, short, usage)
	case IsInt:
		keyInt(flagSet, name, short, usage)
	case IsBool:
		keyBool(flagSet, name, short, usage)
	}

	var flag *pflag.Flag
	if global {
		flag = cmd.PersistentFlags().Lookup(name)
	} else {
		flag = cmd.Flags().Lookup(name)
	}
	return bindFlagToViper(name, flag)
}

// create a string flag
func keyString(flagSet *pflag.FlagSet, name string, short string, usage string) {
	if short == "" {
		flagSet.String(name, "", usage)
	} else {
		flagSet.StringP(name, short, "", usage)
	}
}

// create a boolean flag
func keyBool(flagSet *pflag.FlagSet, name string, short string, usage string) {
	if short == "" {
		flagSet.Bool(name, false, usage)
	} else {
		flagSet.BoolP(name, short, false, usage)
	}
}

// create an integer flag
func keyInt(flagSet *pflag.FlagSet, name string, short string, usage string) {
	if short == "" {
		flagSet.Int(name, 0, usage)
	} else {
		flagSet.IntP(name, short, 0, usage)
	}
}

// bind a cobra flag to a viper key and an environment variable
func bindFlagToViper(name string, flag *pflag.Flag) error {
	err := viper.BindPFlag(name, flag)
	if err != nil {
		return fmt.Errorf("could not bind flag %q to viper key (error=%w)", name, err)
	}

	err = viper.BindEnv(name)
	if err != nil {
		return fmt.Errorf("could not bind flag %q to env key (error=%w)", name, err)
	}
	return nil
}
