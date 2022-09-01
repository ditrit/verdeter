package validators

import (
	"fmt"

	"github.com/ditrit/verdeter/models"
)

// CheckTCPHighPort validate that the port number is a high port
var CheckTCPHighPort = models.Validator{
	Name: "Check if TCP port is high (between 1024 and 65535)",
	Func: func(val interface{}) error {
		intVal, ok := val.(int)
		if !ok {
			return fmt.Errorf("expected an int")
		}
		if intVal >= 1024 && intVal <= 65535 {
			return nil
		}
		return fmt.Errorf("value (%d) is not a TCP high port ", intVal)
	},
}
