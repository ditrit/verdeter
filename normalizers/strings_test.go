package normalizers_test

import (
	"testing"

	"github.com/ditrit/verdeter/normalizers"
	"github.com/stretchr/testify/assert"
)

func TestLowerStrings(t *testing.T) {
	assert.Equal(t, "interface", normalizers.LowerString("INTERFACE"))
	assert.Equal(t, 98, normalizers.LowerString(98))
}

func TestUpperString(t *testing.T) {
	assert.Equal(t, "INTERFACE", normalizers.UpperString("intErface"))
	assert.Equal(t, 98, normalizers.UpperString(98))
}
