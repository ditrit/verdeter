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

func TestValueIn(t *testing.T) {
	set := []int{1, 2, 3, 4}
	validator := ValueIn("integer list", set...)
	assert.NotNil(t, validator)
	assert.Equal(t, "integer list", validator.Name)
	assert.NoError(t, validator.Func(1))
	assert.Error(t, validator.Func("12"))
	assert.Error(t, validator.Func(5))
}
