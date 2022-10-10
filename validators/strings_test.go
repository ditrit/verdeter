package validators_test

import (
	"testing"

	"github.com/ditrit/verdeter/validators"
	"github.com/stretchr/testify/assert"
)

func TestCheckStringNotEmpty(t *testing.T) {
	assert.NoError(t, validators.StringNotEmpty.Func("whatever, hey kid what are you doing ?     "))
	assert.NoError(t, validators.StringNotEmpty.Func("				µ		"), "non empty string with tabs")
	assert.NoError(t, validators.StringNotEmpty.Func("   µ     "), "non empty string with spaces")
	assert.NoError(t, validators.StringNotEmpty.Func("s"))

	assert.Error(t, validators.StringNotEmpty.Func("      "), "empty string with spaces")
	assert.Error(t, validators.StringNotEmpty.Func("			"), "empty string with tabs")
	assert.Error(t, validators.StringNotEmpty.Func(struct{ mohsdhf int }{mohsdhf: 25612564}))
	assert.Error(t, validators.StringNotEmpty.Func(21))
}
