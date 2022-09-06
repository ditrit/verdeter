package validators

import (
	"fmt"
	"strconv"

	"github.com/ditrit/verdeter/models"
)

// CheckTCPHighPort validate that the port number is a high port
var CheckTCPHighPort = models.Validator{
	Name: "Check if TCP port is high (between 1024 and 65535)",
	Func: func(val interface{}) error {
		var intVal int
		var err error
		var ok bool
		intVal, ok = val.(int)
		if !ok {
			stringValue, ok := val.(string)
			intVal, err = strconv.Atoi(stringValue)
			if !ok && err != nil {
				return fmt.Errorf("expected an int")
			}
		}
		if intVal >= 1024 && intVal <= 65535 {
			return nil
		}
		return fmt.Errorf("value (%d) is not a TCP high port ", intVal)
	},
}
