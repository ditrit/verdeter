package validators

import (
	"fmt"

	"github.com/ditrit/verdeter/models"
)

func ValueIn[T comparable](name string, valueSet ...T) models.Validator {
	return models.Validator{
		Name: name,
		Func: func(input interface{}) (err error) {
			inputT, ok := input.(T)
			if !ok {
				var instance T
				return fmt.Errorf("can't convert %v the to %T", input, instance)
			}
			if !contains(inputT, valueSet) {
				return fmt.Errorf("%v is not in %v, please choose a correct value", inputT, valueSet)
			}
			return nil
		},
	}
}

func contains[T comparable](sample T, set []T) bool {
	for _, setValue := range set {
		if setValue == sample {
			return true
		}
	}
	return false
}
