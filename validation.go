package verdeter

import (
	"fmt"

	"github.com/spf13/viper"
)

// Compute and set default values is not set previously
// default values comes from the dynamic set of values
func (vc *VerdeterCommand) ComputeDefaultValues() {
	for key, compute := range vc.computedValue {
		if !viper.IsSet(key) {
			viper.Set(key, compute())
		}
	}
}

// Normalize config values
func (vc *VerdeterCommand) NormalizeValues() {
	for key, normalize := range vc.normalize {
		if viper.IsSet(key) {
			val := viper.Get(key)
			viper.Set(key, normalize(val))
		}
	}
}

// Validate checks if config keys have valid values
func (vc *VerdeterCommand) Validate() error {

	if vc.parentCmd != nil {
		err := vc.parentCmd.Validate()
		if err != nil {
			return fmt.Errorf("(from %q) an error happened while verifying parent command %q : %w", vc.cmd.Name(), vc.parentCmd.cmd.Name(), err)
		}
	}
	vc.ComputeDefaultValues()
	vc.NormalizeValues()

	for key := range vc.isRequired {
		if !viper.IsSet(key) {
			return fmt.Errorf("%q is required", key)
		}
	}

	for key, validator := range vc.isValid {
		valKey := viper.Get(key)
		isSet := viper.IsSet(key)
		if err := validator.Func(valKey); isSet && err != nil {
			return fmt.Errorf("validation %q failed for key %q (ERROR=%w)", validator.Name, key, err)
		}
	}

	for constraintName, constraint := range vc.constraints {
		if !constraint() {
			return fmt.Errorf("constraint %q is not respected", constraintName)
		}
	}

	return nil
}
