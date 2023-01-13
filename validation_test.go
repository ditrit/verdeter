package verdeter_test

import (
	"fmt"
	"testing"

	"github.com/ditrit/verdeter"
	"github.com/stretchr/testify/assert"
)

func TestValidateParent(t *testing.T) {
	cfgParent := verdeter.NewVerdeterCommand(
		"parent",
		"a test command",
		"whatever",
		func(cfg *verdeter.VerdeterCommand, args []string) error {
			fmt.Printf("args=%v", args)
			return nil
		},
	)

	cfgParent.GKey("firstLevel", verdeter.IsStr, "", "first level")
	cfgParent.SetRequired("firstLevel")

	cfgChild := verdeter.NewVerdeterCommand(
		"clid",
		"a test command",
		"whatever",
		func(cfg *verdeter.VerdeterCommand, args []string) error {
			fmt.Printf("args=%v", args)
			return nil
		},
	)
	cfgParent.AddSubCommand(cfgChild)
	assert.Error(t, cfgChild.Validate(true))
}

func TestConstraint(t *testing.T) {
	cfgParent := verdeter.NewVerdeterCommand(
		"parent",
		"a test command",
		"whatever",
		func(cfg *verdeter.VerdeterCommand, args []string) error {
			fmt.Printf("args=%v", args)
			return nil
		},
	)

	cfgParent.SetConstraint("failing constraint", func() bool { return false })
	assert.Error(t, cfgParent.Validate(true))
}
