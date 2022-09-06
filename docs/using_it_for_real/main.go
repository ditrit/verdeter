package main

import (
	"fmt"
	"time"

	"github.com/ditrit/verdeter"
	"github.com/spf13/viper"
)

var verdeterRootCmd = verdeter.NewVerdeterCommand(
	"verdeterapp",
	"verdeterapp print formated time to the terminal",
	"/* Insert a longer description here*/",
	func(cfg *verdeter.VerdeterCommand, args []string) error {
		timeStamp := viper.GetInt("time")
		t := time.Unix(int64(timeStamp), 0)
		fmt.Println(t)

		// no error to return
		return nil
	})

func main() {
	// Initialize the command
	verdeterRootCmd.Initialize()

	viper.Set("config_path", "./conf/")

	// Set a new key named "time" with a shortcut named "t"
	verdeterRootCmd.GKey("time", verdeter.IsInt, "t", "the time")

	// If the value of time is not set, run this function and set "time" to it's output
	verdeterRootCmd.SetComputedValue("time", func() interface{} {
		return time.Now().Unix()
	})

	verdeterRootCmd.Execute()
}
