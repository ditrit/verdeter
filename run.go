package verdeter

import (
	"fmt"

	"github.com/spf13/cobra"
)

// Create the a cobra compatible RunE function
func RunE(cfg *VerdeterCommand, runf func(cfg *VerdeterCommand, args []string) error) func(*cobra.Command, []string) error {
	return func(cobraCmd *cobra.Command, args []string) error {
		return runf(cfg, args)
	}
}

// Create the a cobra compatible PreRunE function.
func preRunCheckE(cfg *VerdeterCommand) func(*cobra.Command, []string) error {
	return func(cobraCmd *cobra.Command, args []string) error {
		if err := cfg.Validate(); err != nil {
			return fmt.Errorf("prerun check for %s failed. (Error=%q))", cfg.cmd.Name(), err.Error())
		} else if len(args) != cfg.nbArgs {
			return fmt.Errorf("prerun check for %s failed. Expected %v args, got %v", cfg.cmd.Name(), cfg.nbArgs, len(args))
		}
		return nil
	}
}
