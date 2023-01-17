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
	listErrors = validateGlobalKeys(verdeterCmd, listErrors)

	// validate local config keys
	if isTargetCommand {
		listErrors = validateLocalKeys(verdeterCmd, listErrors)
	}

	for constraintName, constraint := range verdeterCmd.constraints {
		if !constraint() {
			listErrors = append(listErrors,
				fmt.Errorf("constraint %q is not respected", constraintName),
			)
		}
	}
	if len(listErrors) != 0 {
		printErrors(listErrors)
		return fmt.Errorf("validation failed (%d errors)", len(listErrors))
	}
	return nil
}

// validateLocalKeys: validate global config keys
func validateLocalKeys(verdeterCmd *VerdeterCommand, listErrors []error) []error {
	for _, configKey := range verdeterCmd.localConfigKeys {
		errs := configKey.Validate()
		if errs != nil {
			listErrors = append(listErrors, errs...)
		}
	}
	return listErrors
}

// validateGlobalKeys: validate local config keys
func validateGlobalKeys(verdeterCmd *VerdeterCommand, listErrors []error) []error {
	for _, configKey := range verdeterCmd.globalConfigKeys {
		errs := configKey.Validate()
		if errs != nil {
			listErrors = append(listErrors, errs...)
		}
	}
	return listErrors
}

// printErrors print the error in a clean way.
func printErrors(listErrors []error) {
	fmt.Println("Some errors were collected during initialization:")
	for _, err := range listErrors {
		fmt.Println("-", err.Error())
	}
}
