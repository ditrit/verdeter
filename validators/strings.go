package validators

import (
	"fmt"
	"strings"

	"github.com/ditrit/verdeter/models"
)

// StringNotEmpty validate if a string value is not empty
var StringNotEmpty = models.Validator{
	Name: "String not empty",
	Func: func(input interface{}) error {
		stringVal, ok := input.(string)
		if !ok {
			return fmt.Errorf("expected a string")
		}
		if strings.TrimSpace(stringVal) == "" {
			return fmt.Errorf("empty value is not allowed")
		}
		return nil
	},
}
