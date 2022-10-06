package verdeter

import (
	"testing"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/stretchr/testify/assert"
)

func TestKeyString(t *testing.T) {
	var flagSet *pflag.FlagSet = new(pflag.FlagSet)
	keyString(flagSet, "port", "p", "port description")
	assert.True(t, flagSet.HasFlags())
	assert.Equal(t, "port", flagSet.Lookup("port").Name)
	assert.Equal(t, "p", flagSet.Lookup("port").Shorthand)
	assert.Equal(t, "port description", flagSet.Lookup("port").Usage)
	assert.Equal(t, "string", flagSet.Lookup("port").Value.Type())
}

func TestKeyUint(t *testing.T) {
	var flagSet *pflag.FlagSet = new(pflag.FlagSet)
	keyUint(flagSet, "port", "p", "port description")
	assert.True(t, flagSet.HasFlags())
	assert.Equal(t, "port", flagSet.Lookup("port").Name)
	assert.Equal(t, "p", flagSet.Lookup("port").Shorthand)
	assert.Equal(t, "uint", flagSet.Lookup("port").Value.Type())
}

func TestKeyInt(t *testing.T) {
	var flagSet *pflag.FlagSet = new(pflag.FlagSet)
	keyInt(flagSet, "port", "p", "port description")
	assert.True(t, flagSet.HasFlags())
	assert.Equal(t, "port", flagSet.Lookup("port").Name)
	assert.Equal(t, "p", flagSet.Lookup("port").Shorthand)
	assert.Equal(t, "int", flagSet.Lookup("port").Value.Type())
}

func TestKeyBool(t *testing.T) {
	var flagSet *pflag.FlagSet = new(pflag.FlagSet)
	keyBool(flagSet, "port", "p", "port description")
	assert.True(t, flagSet.HasFlags())
	assert.Equal(t, "port", flagSet.Lookup("port").Name)
	assert.Equal(t, "p", flagSet.Lookup("port").Shorthand)
	assert.Equal(t, "bool", flagSet.Lookup("port").Value.Type())
}

func TestKeyGlobalTrue(t *testing.T) {
	// testing if the flag end up in the global flags
	var cmd *cobra.Command = new(cobra.Command)
	Key(cmd, "port", IsStr, "p", "port usage", true)
	assert.True(t, cmd.PersistentFlags().HasFlags())
	assert.False(t, cmd.Flags().HasFlags())
}

func TestKeyGlobalFalse(t *testing.T) {
	// testing if the flag end in the local flags
	var cmd *cobra.Command = new(cobra.Command)
	Key(cmd, "port", IsStr, "p", "port usage", false)
	assert.False(t, cmd.PersistentFlags().HasFlags())
	assert.True(t, cmd.Flags().HasFlags())
}

func TestKey(t *testing.T) {
	var cmd *cobra.Command = new(cobra.Command)
	Key(cmd, "port", IsStr, "p", "port usage", true)
	assert.Equal(t, "port", cmd.PersistentFlags().Lookup("port").Name)
	assert.Equal(t, "p", cmd.PersistentFlags().Lookup("port").Shorthand)
	assert.Equal(t, "port usage", cmd.PersistentFlags().Lookup("port").Usage)
}

func TestKeyIsStr(t *testing.T) {
	var cmd *cobra.Command = new(cobra.Command)
	Key(cmd, "port", IsStr, "p", "port usage", false)
	println(cmd.PersistentFlags())
	assert.Equal(t, "string", cmd.Flags().Lookup("port").Value.Type())
}

func TestKeyIsBool(t *testing.T) {
	var cmd *cobra.Command = new(cobra.Command)
	Key(cmd, "port", IsBool, "p", "port usage", false)
	assert.Equal(t, "bool", cmd.Flags().Lookup("port").Value.Type())
}

func TestKeyIsInt(t *testing.T) {
	var cmd *cobra.Command = new(cobra.Command)
	Key(cmd, "port", IsInt, "p", "port usage", false)
	assert.Equal(t, "int", cmd.Flags().Lookup("port").Value.Type())
}

func TestKeyIsUint(t *testing.T) {
	var cmd *cobra.Command = new(cobra.Command)
	Key(cmd, "port", IsUint, "p", "port usage", false)
	assert.Equal(t, "uint", cmd.Flags().Lookup("port").Value.Type())
}
