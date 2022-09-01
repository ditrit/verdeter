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
		println(args)
		println("value for \"key=offset\":", viper.GetInt("offset"))
		println("value for \"key=voila\":", viper.GetString("voila"))
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
	myCommand.Initialize()

	myCommand.GKey("config_path", verdeter.IsStr, "", "Path to the config directory/file")
	myCommand.GKey("rootnode", verdeter.IsInt, "", "is root node")
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
