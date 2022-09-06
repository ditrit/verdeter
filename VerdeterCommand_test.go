package verdeter_test

import (
	"fmt"
	"testing"

	"github.com/ditrit/verdeter"
	"github.com/stretchr/testify/assert"
)

// testing contructor of VerdeterCommand
func TestNewVerdeterCommand(t *testing.T) {
	verdeterCmd := verdeter.NewVerdeterCommand(
		"test",
		"test short description",
		"test long description",
		func(verdeterCmd *verdeter.VerdeterCommand, args []string) error {
			fmt.Println("hello mom")
			return nil
		},
	)
	assert.NotNil(t, verdeterCmd, "func NewVerdeterCommand should not return something nil")
}
