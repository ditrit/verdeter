package models

// A validator is used when it's needed to validate an input from verdeter,
// whether the input comes from the user of the app or a default value
type Validator struct {
	// The name of the validator
	// It's printed when the Validation Function returns an error
	Name string

	// The validation function, it returns an error when the input is not valid
	Func ValidationFunction
}
