package main

import (
	"github.com/ditrit/verdeter"
	"github.com/spf13/viper"
)

var myCommand = verdeter.NewVerdeterCommand(
	"math", // the name of the command

	"math app is a test application foir verdeter",

	``,

	func(cfg *verdeter.VerdeterCommand, args []string) error {
		println("The root command does nothing but print some config keys")
		println("args:", args)
		println("value for \"key=some.config\":", viper.GetUint("some.config"))
		println("value for \"key=rootnode\":", viper.GetInt("rootnode"))
		return nil
	})

var add = verdeter.NewVerdeterCommand(
	"add",
	"print the result of --int1 + --int2",
	``,
	func(cfg *verdeter.VerdeterCommand, args []string) error {
		println("print the result of --int1 + --float2 (or -i + -f)")
		println("result", viper.GetInt("int2")+viper.GetInt("int1"))
		return nil
	})

func init() {
	myCommand.GKey("config_path", verdeter.IsStr, "", "Path to the config directory/file")
	myCommand.GKey("rootnode", verdeter.IsInt, "", "is root node")
	myCommand.GKey("some.config", verdeter.IsUint, "", "is root node")
	myCommand.SetRequired("some.config")
	myCommand.SetNormalize("config_path", func(val interface{}) interface{} {
		strval, ok := val.(string)
		if ok && strval != "" {
			lastChar := strval[len(strval)-1:]
			if lastChar != "/" {
				return strval + "/"
			}
			return strval
		}
		return nil
	})
	add.LKey("int1", verdeter.IsInt, "", "Integer 1")
	add.LKey("int2", verdeter.IsInt, "", "Integer 2")
	add.SetRequired("int1")
	add.SetRequired("int2")
	myCommand.AddSubCommand(add)
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	myCommand.Execute()
}

func main() {
	Execute()
}
