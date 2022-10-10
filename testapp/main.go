package main

import (
	"github.com/ditrit/verdeter"
	"github.com/spf13/viper"
)

var myCommand = verdeter.NewVerdeterCommand(
	"myapp", // the name of the command

	"myapp is a verdeter com",

	``,

	func(cfg *verdeter.VerdeterCommand, args []string) error {
		println("args:", args)
		println("value for \"key=some.config\":", viper.GetUint("some.config"))
		println("value for \"key=rootnode\":", viper.GetInt("rootnode"))

		// no error to return
		return nil
	})

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	myCommand.Execute()
}

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

}

func main() {
	Execute()
}
