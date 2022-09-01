package verdeter

import (
	"fmt"
	"strconv"

	"github.com/spf13/viper"
)

// Return the offset
// if the viper key "offset" in no set, it returns a 0
func GetOffset() int {
	return viper.GetInt("offset")
}

// Bind the environment variable for the offset key and set the offset as an integer
func initOffset(vc *VerdeterCommand) {
	viper.SetEnvPrefix(vc.GetAppName())
	viper.BindEnv("offset")
	viper.SetDefault("offset", 0)
	valStr := viper.GetString("offset")
	var offset int64
	offset, err := strconv.ParseInt(valStr, 10, 0)
	if err != nil {
		panic(fmt.Errorf("offset is not an integer (got %v)", valStr))
	}
	viper.Set("offset", int(offset))
}
