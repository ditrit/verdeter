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
func (configKey *ConfigKey) Validate() error {
	configKey.ComputeDefaultValue()
	configKey.Normalize()

	err := configKey.CheckRequired()
	if err != nil {
		return err
	}

	err = configKey.CheckValidators()
	if err != nil {
		return err
	}

	return nil
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
func (configKey *ConfigKey) CheckValidators() error {
	for _, validator := range configKey.validators {
		if !viper.IsSet(configKey.Name) {
			continue
		}
		valKey := viper.Get(configKey.Name)
		err := validator.Func(valKey)
		if err != nil {
			return fmt.Errorf("validation %q failed for key %q (ERROR=%w)", validator.Name, configKey.Name, err)
		}
	}
	return nil
}
