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
		err := initConfig(cfg)
		if err != nil {
			return err
		}

		if err := cfg.Validate(true); err != nil {
			return err
		} else if len(args) != cfg.nbArgs {
			return fmt.Errorf("expected %v args, got %v", cfg.nbArgs, len(args))
		}
		return nil
	}
}
