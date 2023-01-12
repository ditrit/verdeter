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
	listErrors := make([]error, 0)

	// validate global config keys
	for _, configKey := range verdeterCmd.globalConfigKeys {
		errs := configKey.Validate()
		if errs != nil {
			listErrors = append(listErrors, errs...)
		}
	}

	// validate local config keys
	if isTargetCommand {
		for _, configKey := range verdeterCmd.localConfigKeys {
			errs := configKey.Validate()
			if errs != nil {
				listErrors = append(listErrors, errs...)
			}
		}
	}

	for constraintName, constraint := range verdeterCmd.constraints {
		if !constraint() {
			listErrors = append(listErrors,
				fmt.Errorf("constraint %q is not respected", constraintName),
			)
		}
	}
	printErrors(listErrors)
	if len(listErrors) == 0 {
		return nil
	}
	return fmt.Errorf("validation failed")
}

func printErrors(listErrors []error) {
	for _, err := range listErrors {
		fmt.Println(err.Error())
	}
}
