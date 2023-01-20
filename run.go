package verdeter

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// buildPreRunCheckE create a cobra compatible PreRunE function.
func (verdeterCmd *VerdeterCommand) buildPreRunCheckE(preRunE func(cobraCmd *cobra.Command, args []string) error) func(*cobra.Command, []string) error {
	if preRunE == nil {
		return nil
	}
	return func(cobraCmd *cobra.Command, args []string) error {
		err := verdeterCmd.initAndValidate()
		if err != nil {
			return err
		}
		if preRunE == nil {
			return nil
		}
		return preRunE(cobraCmd, args)
	}
}

// initAndValidate initialize the command
func (verdeterCmd *VerdeterCommand) initAndValidate() error {
	err := initConfig(verdeterCmd)
	if err != nil {
		return err
	}
	if err := verdeterCmd.Validate(true); err != nil {
		return fmt.Errorf("prerun check for %s failed. (Error=%q))", verdeterCmd.cmd.Name(), err.Error())
	}
	return nil
}

// buildPreRunCheck create a cobra compatible PreRun function.
func (verdeterCmd *VerdeterCommand) buildPreRunCheck(preRun func(cobraCmd *cobra.Command, args []string)) func(*cobra.Command, []string) {

	return func(cobraCmd *cobra.Command, args []string) {
		err := verdeterCmd.initAndValidate()
		if err != nil {
			fmt.Printf("An error was returned: %s\n", err.Error())
			os.Exit(1)
		}
		if preRun == nil {
			return
		}
		preRun(cobraCmd, args)
	}
}
