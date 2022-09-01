package verdeter_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/ditrit/verdeter"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestNormalUse(t *testing.T) {
	os.Setenv("TEST_OFFSET", "1")
	os.Setenv("TEST_1_ENVKEY", "envkeyvalue")
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
	cfg.SetDefault("offset", 1)

	cfg.GKey("config_path", verdeter.IsStr, "", "path to the config directory")
	cfg.SetNormalize("config_path", func(val interface{}) interface{} {
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

	assert.Equal(t, "test", cfg.GetAppName(), "should return the name of the app")
	assert.Equal(t, cfg.GetInstanceName(), "test_1", "should return the name of the app")
	cfg.LKey("superkey", verdeter.IsInt, "", "test key in fixture dir")
	cfg.LKey("envkey", verdeter.IsStr, "", "test env key")
	cfg.SetDefault("superkey", -5)

	cfg.Execute()

	assert.Equal(t, 125, viper.GetInt("superkey"))
	assert.Equal(t, "envkeyvalue", viper.GetString("envkey"), "should be equals")
	assert.NoError(t, cfg.Validate(), "shouldn't ")

}
