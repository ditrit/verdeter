package verdeter

import (
	"fmt"

	"github.com/ditrit/verdeter/models"
	"github.com/spf13/viper"
)

// Represent a Config Key
type ConfigKey struct {
	// the name of the config key
	Name          string
	validators    []models.Validator
	required      bool
	computedValue models.DefaultValueFunction
	normalizeFunc models.NormalizationFunction
}

// Validate the configkey
//
// 1. Run the dynamic default function
// 2. Normalize the value
// 3. Check the required constraint
// 4. Check the validators
func (configKey *ConfigKey) Validate() []error {
	configKey.ComputeDefaultValue()
	configKey.Normalize()
	result := make([]error, 0)

	err := configKey.CheckRequired()
	if err != nil {
		return append(result, err)
	}

	errs := configKey.CheckValidators()
	if errs != nil {
		result = append(result, errs...)
	}

	return result
}

// Normalize the config key using the normalization function if provided
func (configKey *ConfigKey) Normalize() {
	if viper.IsSet(configKey.Name) && configKey.normalizeFunc != nil {
		val := viper.Get(configKey.Name)
		viper.Set(configKey.Name, configKey.normalizeFunc(val))
	}
}

// Compute the default value of the config key using the DefaultValueFunction function if provided
func (configKey *ConfigKey) ComputeDefaultValue() {
	if !viper.IsSet(configKey.Name) && configKey.computedValue != nil {
		viper.Set(configKey.Name, configKey.computedValue())
	}
}

// Check if the value is provided if the config key is required
// Return an error on failure
func (configKey *ConfigKey) CheckRequired() error {
	if !viper.IsSet(configKey.Name) && configKey.required {
		return fmt.Errorf("%q is required", configKey.Name)
	}
	return nil
}

// Return an error on validation failure of one of the validator.
//
// Return on first failure, the remaining validator are not ran.
func (configKey *ConfigKey) CheckValidators() []error {
	result := make([]error, 0)
	for _, validator := range configKey.validators {
		if !viper.IsSet(configKey.Name) {
			continue
		}
		valKey := viper.Get(configKey.Name)
		err := validator.Func(valKey)
		if err != nil {
			result = append(result,
				fmt.Errorf("validator %q failed for key %q: %q", validator.Name, configKey.Name, err.Error()),
			)
		}
	}
	return result
}
