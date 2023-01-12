package validators

import (
	"fmt"
	"net/url"

	"github.com/ditrit/verdeter/models"
)

// URLValidator validate that string provided is an url
var URLValidator = models.Validator{
	Name: "Url Validator",
	Func: func(val interface{}) error {
		valString, ok := val.(string)
		if !ok {
			return fmt.Errorf("'%v' is not a even a string, can't be an url", val)
		}
		// See https://cs.opensource.google/go/go/+/refs/tags/go1.19.4:src/net/url/url_test.go;l=672
		// for tests
		_, err := url.ParseRequestURI(valString)
		if err != nil {
			return fmt.Errorf("failed to parse %q as an url", valString)
		}
		return nil
	},
}
