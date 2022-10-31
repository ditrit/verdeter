package verdeter

import (
	"fmt"
)

// Validate checks if config keys have valid values
func (verdeterCmd *VerdeterCommand) Validate(isTargetCommand bool) error {
	// Validate parent command
	if verdeterCmd.parentCmd != nil {
		err := verdeterCmd.parentCmd.Validate(false)
		if err != nil {
			return fmt.Errorf("(from %q) an error happened while verifying parent command %q : %w", verdeterCmd.cmd.Name(), verdeterCmd.parentCmd.cmd.Name(), err)
		}
	}

	// validate global config keys
	for _, configKey := range verdeterCmd.globalConfigKeys {
		err := configKey.Validate()
		if err != nil {
			return err
		}
	}

	// validate local config keys
	if isTargetCommand {
		for _, configKey := range verdeterCmd.localConfigKeys {
			err := configKey.Validate()
			if err != nil {
				return err
			}
		}
	}

	for constraintName, constraint := range verdeterCmd.constraints {
		if !constraint() {
			return fmt.Errorf("constraint %q is not respected", constraintName)
		}
	}

	return nil
}
