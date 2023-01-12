package validators_test

import (
	"testing"

	"github.com/ditrit/verdeter/validators"
	"github.com/stretchr/testify/assert"
)

func TestUrlParseable(t *testing.T) {
	assert.Equal(t, "Url Validator", validators.URLValidator.Name)
	assert.Error(t, validators.URLValidator.Func("qsdqsdqs"))
	assert.Error(t, validators.URLValidator.Func(1521645))
	assert.NoError(t, validators.URLValidator.Func("https://github.com/ditrit/verdeter"))
	assert.NoError(t, validators.URLValidator.Func("gemini://qsd.com/qsd"))
}
