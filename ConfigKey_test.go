package verdeter

import (
	"errors"
	"testing"

	"github.com/ditrit/verdeter/models"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestConfigKeyComputedValue(t *testing.T) {
	viper.Reset()
	cf := &ConfigKey{
		Name: "test.ConfigKey",
		computedValue: func() (dynamicDefault interface{}) {
			return 2
		},
	}

	cf.ComputeDefaultValue()
	assert.Equal(t, 2, viper.GetInt("test.ConfigKey"))
	viper.Set("test.ConfigKey", 1)
	cf.ComputeDefaultValue()
	assert.Equal(t, 1, viper.GetInt("test.ConfigKey"))
}

func TestConfigKeyCheckRequired(t *testing.T) {
	viper.Reset()
	cf := &ConfigKey{
		Name:     "test.ConfigKey",
		required: false,
	}

	assert.NoError(t, cf.CheckRequired())
	cf.required = true
	assert.Error(t, cf.CheckRequired())
	viper.Set("test.ConfigKey", 1)
	assert.NoError(t, cf.CheckRequired())
}

func TestConfigKeyNormalize(t *testing.T) {
	viper.Reset()
	cf := &ConfigKey{
		Name: "test.ConfigKey",
		normalizeFunc: func(input interface{}) (output interface{}) {
			intVal := input.(int)
			return intVal + 1
		},
	}

	viper.Set("test.ConfigKey", 1)
	assert.Equal(t, 1, viper.GetInt("test.ConfigKey"))
	cf.Normalize()
	assert.Equal(t, 2, viper.GetInt("test.ConfigKey"))
}

func TestConfigKeyValidators(t *testing.T) {
	viper.Reset()
	cf := &ConfigKey{
		Name: "test.ConfigKey",
		validators: []models.Validator{
			{
				Name: "greater than 1",
				Func: func(input interface{}) error {
					intVal := input.(int)
					if intVal < 1 {
						return errors.New("greater than 1")
					}
					return nil
				},
			}, {
				Name: "greater than 5",
				Func: func(input interface{}) error {
					intVal := input.(int)
					if intVal < 5 {
						return errors.New("greater than 5")
					}
					return nil
				},
			},
		},
	}

	viper.Set("test.ConfigKey", 0)
	assert.Len(t, cf.CheckValidators(), 2)
	viper.Set("test.ConfigKey", 2)
	assert.Len(t, cf.CheckValidators(), 1)
	viper.Set("test.ConfigKey", 12)
	assert.Len(t, cf.CheckValidators(), 0)
}

func TestConfigKeyValidate(t *testing.T) {
	viper.Reset()
	cf := &ConfigKey{
		Name: "test.ConfigKey",
		validators: []models.Validator{
			{
				Name: "greater than 1",
				Func: func(input interface{}) error {
					intVal := input.(int)
					if intVal <= 1 {
						return errors.New("greater than 1")
					}
					return nil
				},
			}, {
				Name: "greater than 5",
				Func: func(input interface{}) error {
					intVal := input.(int)

					if intVal <= 5 {
						return errors.New("greater than 5")
					}
					return nil
				},
			},
		},
		required: false,
		computedValue: func() (dynamicDefault interface{}) {
			return 50
		},
		normalizeFunc: func(input interface{}) (output interface{}) {
			intVal := input.(int)
			return intVal + 1
		},
	}
	viper.Reset()
	assert.Len(t, cf.Validate(), 0)
	assert.Equal(t, 51, viper.GetInt("test.ConfigKey"))

	viper.Reset()
	viper.Set("test.ConfigKey", -1)
	assert.Len(t, cf.Validate(), 2)
	viper.Reset()

	viper.Set("test.ConfigKey", 1)
	assert.Len(t, cf.Validate(), 1)

	viper.Set("test.ConfigKey", 11)
	assert.Len(t, cf.Validate(), 0)
}
