package main

import (
	"errors"
	"fmt"

	"github.com/ditrit/verdeter"
	"github.com/ditrit/verdeter/models"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var myCommand = verdeter.BuildVerdeterCommand(verdeter.VerdeterConfig{
	Use:         "math",
	Long:        "math app is a test application for verdeter",
	Annotations: map[string]string{},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("The root command does nothing but print the config key 'organisation.name'")
		fmt.Println("value for \"key=organisation.name\":", viper.GetString("organisation.name"))
	},
	RunE: nil,
})

var add = verdeter.NewVerdeterCommand(
	"add",
	"print the result of --int1 + --int2",
	``,
	func(cfg *verdeter.VerdeterCommand, args []string) error {
		println("print the result of --int1 + --int2")
		println("result:", viper.GetInt("int2")+viper.GetInt("int1"))
		return nil
	})

func init() {
	myCommand.GKey("config_path", verdeter.IsStr, "", "Path to the config directory/file")
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
	myCommand.GKey("organisation.name", verdeter.IsStr, "", "An organisation name")
	myCommand.SetRequired("organisation.name")

	add.LKey("int1", verdeter.IsInt, "", "Integer 1")
	add.LKey("int2", verdeter.IsInt, "", "Integer 2")
	add.SetRequired("int1")
	add.SetRequired("int2")

	over0 := models.Validator{
		Name: "Superior to 0",
		Func: func(input interface{}) (err error) {
			intVal, ok := input.(int)
			if !ok {
				return errors.New("rong type")
			}
			if intVal < 0 {
				return errors.New("under1")
			}

			return nil
		},
	}
	add.AddValidator("int1", over0)
	add.AddValidator("int2", over0)

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
