package validators_test

import (
	"testing"

	"github.com/ditrit/verdeter/validators"
	"github.com/stretchr/testify/assert"
)

func TestCheckStringNotEmpty(t *testing.T) {
	testCases := []struct {
		desc  string
		in    any
		valid bool
	}{
		{
			desc:  "valid: not empty",
			in:    "whatever, hey kid what are you doing ?     ",
			valid: true,
		}, {
			desc:  "invalid: only spaces",
			in:    "     ",
			valid: false,
		}, {
			desc:  "invalid: only tabs",
			in:    "			",
			valid: false,
		}, {
			desc:  "valid: only one char",
			in:    "s",
			valid: true,
		}, {
			desc:  "valid: one char and spaces",
			in:    "                µ        ",
			valid: true,
		}, {
			desc:  "valid: one char and tabs",
			in:    "				µ		",
			valid: true,
		}, {
			desc:  "valid: one char and spaces and tabs",
			in:    "	  	 		 µ	 	",
			valid: true,
		}, {
			desc:  "invalid: wrong input type",
			in:    21,
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
				assert.NoError(t, validators.StringNotEmpty.Func(tc.in), "should not return an error")
			} else {
				assert.Error(t, validators.StringNotEmpty.Func(tc.in), "should return an error")
			}

		})
	}
}
