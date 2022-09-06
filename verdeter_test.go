package verdeter_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/ditrit/verdeter"
	"github.com/ditrit/verdeter/validators"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestNormalUse(t *testing.T) {
	// set an env key for later
	os.Setenv("TEST_ENVKEY", "envkeyvalue")
	cfg := verdeter.NewVerdeterCommand(
		"test",
		"a test command",
		"whatever",
		func(cfg *verdeter.VerdeterCommand, args []string) error {
			fmt.Printf("args=%v", args)
			return nil
		},
	)
	cfg.Initialize()

	// Test the conputed values works
	cfg.GKey("computed", verdeter.IsInt, "", "a computed value")
	cfg.SetComputedValue("computed", func() interface{} {
		return 1234
	})

	// Test the config file handling system
	cfg.GKey("config_path", verdeter.IsStr, "", "path to the config directory")
	cfg.SetNormalize("config_path",
		func(val interface{}) interface{} {
			strval, ok := val.(string)
			if ok {
				if strval != "" {
					lastChar := strval[len(strval)-1:]
					if lastChar != "/" {
						return strval + "/"
					}
					return strval
				}
			}
			return nil
		})
	cfg.SetDefault("config_path", "./fixtures/1/")
	// Test contraints
	cfg.SetConstraint("contraintname", func() bool {
		return true
	})
	cfg.LKey("superkey", verdeter.IsInt, "", "test key in fixture dir")
	cfg.LKey("envkey", verdeter.IsStr, "", "test env key")
	cfg.SetRequired("envkey")
	cfg.SetValidator("envkey", validators.StringNotEmpty)
	cfg.SetDefault("superkey", -5)
	cfg.SetNbArgs(0)

	sub := verdeter.NewVerdeterCommand(
		"sub",
		"a sub command",
		"whatever",
		func(cfg *verdeter.VerdeterCommand, args []string) error {
			fmt.Printf("args=%v", args)
			return nil
		},
	)
	cfg.AddSubCommand(sub)

	cfg.Execute()

	assert.Equal(t, "test", cfg.GetAppName(), "should return the name of the app")
	assert.Equal(t, 125, viper.GetInt("superkey"))
	assert.Equal(t, 1234, viper.GetInt("computed"))
	assert.Equal(t, "envkeyvalue", viper.GetString("envkey"), "should be equals")
	assert.NoError(t, cfg.Validate(), "shouldn't ")

}

func TestForceConfigFile(t *testing.T) {
	cfg := verdeter.NewVerdeterCommand(
		"test",
		"a test command",
		"whatever",
		func(cfg *verdeter.VerdeterCommand, args []string) error {
			fmt.Printf("args=%v", args)
			return nil
		},
	)
	cfg.Initialize()

	cfg.GKey("config_path", verdeter.IsStr, "", "path to the config directory")

	cfg.SetDefault("config_path", "./fixtures/1/")
	viper.Set("config_path", "/truc/biule.yml")
	assert.Panics(t, func() { cfg.Execute() }, "should panic")

}
