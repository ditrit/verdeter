package validators

import (
	"fmt"

	"github.com/ditrit/verdeter/models"
)

// Return a Validator that check if the value received in the validation process is contained in the array of authorized values.
func AuthorizedValues[T comparable](name string, authorizedValues ...T) models.Validator {
	return models.Validator{
		Name: name,
		Func: func(input interface{}) (err error) {
			inputAsT, ok := input.(T)
			if !ok {
				var instance T
				return fmt.Errorf("can't convert %v the to %T", input, instance)
			}
			if !contains(inputAsT, authorizedValues) {
				return fmt.Errorf("%v is not in %v, please choose a correct value", inputAsT, authorizedValues)
			}
			return nil
		},
	}
}

// return true is the sample is found in the set
func contains[T comparable](sample T, set []T) bool {
	for _, setValue := range set {
		if setValue == sample {
			return true
		}
	}
	return false
}
