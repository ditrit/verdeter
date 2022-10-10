package validators

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContains(t *testing.T) {
	set := []int{1, 2, 3, 4}
	assert.True(t, contains(1, set))
	assert.False(t, contains(5, set))
}

func TestAuthorizedValuesInt(t *testing.T) {
	set := []int{1, 2, 3, 4}
	validator := AuthorizedValues("integer list", set...)
	assert.NotNil(t, validator)
	assert.Equal(t, "integer list", validator.Name)
	assert.NoError(t, validator.Func(1))
	assert.Error(t, validator.Func("12"))
	assert.Error(t, validator.Func(5))
}

func TestAuthorizedValuesString(t *testing.T) {
	set := []string{"1", "2", "3", "4"}
	validator := AuthorizedValues("string list", set...)
	assert.NotNil(t, validator)
	assert.Equal(t, "string list", validator.Name)
	assert.NoError(t, validator.Func("1"))
	assert.Error(t, validator.Func("12"))
	assert.Error(t, validator.Func(4))
}

func TestAuthorizedValuesUint(t *testing.T) {
	set := []uint{1, 2, 3, 4}
	validator := AuthorizedValues("uint list", set...)
	assert.NotNil(t, validator)
	assert.Equal(t, "uint list", validator.Name)
	assert.NoError(t, validator.Func(uint(1)))
	assert.Error(t, validator.Func("12"))
	assert.Error(t, validator.Func(uint(5)))
}
