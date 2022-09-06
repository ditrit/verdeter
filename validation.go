package verdeter

import (
	"fmt"

	"github.com/spf13/viper"
)

// Compute and set default values is not set previously
// default values comes from the dynamic set of values
func (verdeterCmd *VerdeterCommand) ComputeDefaultValues() {
	for key, compute := range verdeterCmd.computedValue {
		if !viper.IsSet(key) {
			viper.Set(key, compute())
		}
	}
}

// Normalize config values
func (verdeterCmd *VerdeterCommand) NormalizeValues() {
	for key, normalize := range verdeterCmd.normalize {
		if viper.IsSet(key) {
			val := viper.Get(key)
			viper.Set(key, normalize(val))
		}
	}
}

// Validate checks if config keys have valid values
func (verdeterCmd *VerdeterCommand) Validate() error {

	if verdeterCmd.parentCmd != nil {
		err := verdeterCmd.parentCmd.Validate()
		if err != nil {
			return fmt.Errorf("(from %q) an error happened while verifying parent command %q : %w", verdeterCmd.cmd.Name(), verdeterCmd.parentCmd.cmd.Name(), err)
		}
	}
	verdeterCmd.ComputeDefaultValues()
	verdeterCmd.NormalizeValues()

	for key := range verdeterCmd.isRequired {
		if !viper.IsSet(key) {
			return fmt.Errorf("%q is required", key)
		}
	}

	for key, validator := range verdeterCmd.isValid {
		valKey := viper.Get(key)
		isSet := viper.IsSet(key)
		if err := validator.Func(valKey); isSet && err != nil {
			return fmt.Errorf("validation %q failed for key %q (ERROR=%w)", validator.Name, key, err)
		}
	}

	for constraintName, constraint := range verdeterCmd.constraints {
		if !constraint() {
			return fmt.Errorf("constraint %q is not respected", constraintName)
		}
	}

	return nil
}