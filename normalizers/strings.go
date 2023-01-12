package normalizers

import (
	"strings"

	"github.com/ditrit/verdeter/models"
)

// LowerString is a NormalizationFunction that lower strings (don't do anything to other types)
var LowerString models.NormalizationFunction = func(input interface{}) (output interface{}) {
	inputStr, ok := input.(string)
	if !ok {
		return input
	}
	return strings.ToLower(inputStr)
}

// UpperString is a NormalizationFunction that upper strings (don't do anything to other types)
var UpperString models.NormalizationFunction = func(input interface{}) (output interface{}) {
	inputStr, ok := input.(string)
	if !ok {
		return input
	}
	return strings.ToUpper(inputStr)
}
