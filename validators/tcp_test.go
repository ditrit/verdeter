package validators_test

import (
	"testing"

	"github.com/ditrit/verdeter/validators"
	"github.com/stretchr/testify/assert"
)

func TestCheckTCPHighPort(t *testing.T) {
	testCases := []struct {
		desc  string
		in    any
		valid bool
	}{
		{
			desc:  "invalid: too low",
			in:    1023,
			valid: false,
		}, {
			desc:  "invalid: too high",
			in:    65536,
			valid: false,
		}, {
			desc:  "valid: middle of the range",
			in:    32768,
			valid: true,
		}, {
			desc:  "valid: low border",
			in:    1024,
			valid: true,
		}, {
			desc:  "valid: high border",
			in:    65535,
			valid: true,
		}, {
			desc:  "invalid: wrong input type",
			in:    "whatever",
			valid: false,
		}, {
			desc:  "invalid: wrong input type",
			in:    struct{ mohsdhf int }{mohsdhf: 25612564},
			valid: false,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			if tc.valid {
				assert.NoError(t, validators.CheckTCPHighPort.Func(tc.in), "should not return an error")
			} else {
				assert.Error(t, validators.CheckTCPHighPort.Func(tc.in), "should return an error")
			}

		})
	}
}
