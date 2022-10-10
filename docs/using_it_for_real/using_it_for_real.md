# Using Verdeter, for real

We will create an app that will take a unix timestamp as an input for a config key and print a formated version to the standard output.

Let's define our root command with the callback
```go
var verdeterRootCmd = verdeter.NewVerdeterCommand(
	"verdeterapp",
	"verdeterapp print formated time to the terminal",
	"/* Insert a longer description here*/",
	func(cfg *verdeter.VerdeterCommand, args []string) error {
		timeStamp := viper.GetInt("time")
		t := time.Unix(int64(timeStamp), 0)
		fmt.Println(t)

		// no error to return
		return nil
})
```

Then set the config_path to `./conf/vertederapp.yml`
```go
verdeterRootCmd.SetDefault("config_path", "./conf/")
```

Then add a consig key 

```go
// Set a new key named "time" with a shortcut named "t"
verdeterRootCmd.GKey("time", verdeter.IsInt, "t", "the input unix time")
```

The we set a dynamic default value to `time.Now().Unix()`. That way if the config key 'time" is not present in the flag, in the environment variables or in the config file, "time" will take the value of the current time.


```go
// If the value of time is not set, run this function and set "time" to it's output
verdeterRootCmd.SetComputedValue("time", func() interface{} {
    return time.Now().Unix()
})
```

The we set a dynamic default value to `time.Now().Unix()`. That way if the config key 'time" is not present in the flag, in the environment variables or in the config file, "time" will take the value of the current time.


```go
// If the value of time is not set, run this function and set "time" to it's output
verdeterRootCmd.SetComputedValue("time", func() interface{} {
    return time.Now().Unix()
})
```
Let's write a config file

```yml
# ./conf/verdeterapp.yml
time: 1661865582
```
***The code in full:***
```go
package main

import (
	"fmt"
	"time"

	"github.com/ditrit/verdeter"
	"github.com/spf13/viper"
)

var verdeterRootCmd = verdeter.NewVerdeterCommand(
	"verdeterapp",
	"verdeterapp print formated time to the terminal",
	"/* Insert a longer description here*/",
	func(cfg *verdeter.VerdeterCommand, args []string) error {
		timeStamp := viper.GetInt("time")
		t := time.Unix(int64(timeStamp), 0)
		fmt.Println(t)

		// no error to return
		return nil
	})

func main() {
	// Initialize the command
	viper.Set("config_path", "./conf/")

	// Set a new key named "time" with a shortcut named "t"
	verdeterRootCmd.GKey("time", verdeter.IsInt, "t", "the time")

	// If the value of time is not set, run this function and set "time" to it's output
	verdeterRootCmd.SetComputedValue("time", func() interface{} {
		return time.Now().Unix()
	})

	verdeterRootCmd.Execute()
}

```


## Executing the program

> the "help" command is available
 ```
$ verdeterapp help
Usage:
    verdeterapp [flags]

Flags:
    -h, --help       help for verdeterapp
    -t, --time int   the time
 ```

**`time` can now be set:**
- with an environment variable `VERDETERAPP_TIME` 
- with a flag `--time` (or `-t`)
- with a value set in a config 



"time" is set from the flag

```bash
> go build -o verdeterapp
> ./verdeterapp --time 1661865571
Tue Aug 30 2022 13:19:31 GMT+0000
```

"time" is set from the ENV variable 

```bash
> go build -o verdeterapp
> VERDETERAPP_TIME=1661865555 ./verdeterapp
Tue Aug 30 2022 13:19:15 GMT+0000
```

"time" is set from the config file 

```bash
> go build -o verdeterapp
> ./verdeterapp
Tue Aug 30 2022 13:19:42 GMT+0000
```

In this example, verdeter allowed us to build an app that can use configuration from 3 main sources (*flags, environment variables and config file*) easily.