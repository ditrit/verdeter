package models

// Returns a normalized version of the input.
//
// For examples of implementation, please head to the normalizators package
type NormalizationFunction func(input interface{}) (output interface{})

// Returns an error of the input is not valid
//
// For examples of implementation, please head to the validators package
type ValidationFunction func(input interface{}) (err error)

// Function that return a default value, this value may be dynamic
type DefaultValueFunction func() (dynamicDefault interface{})

// validate constraints independent from the inputs
//
// (As an example, we could imagine a cli subcommand that could only be run on Unix based systems.
// In this case, a constraint function allow the developper to restrict the use of a specific command and it's subcommands to Unix based systems)
type ConstraintFunction func() bool
