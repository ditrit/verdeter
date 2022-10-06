# Verdeter

Verdeter is a library to write configuration easily with cobra and viper for distributed applications. Verdeter bring the power of cobra and viper in a single library. 

It should be consider as a wrapper for cobra and viper that allow developers to code faster.

> The api is susceptible to change at any point in time until the v1 is released.

Verdeter allow developers to bind a posix complient flag, an environment variable and a variable in a config file to a viper key with a single line of code. 
Verdeter also comes with extra features such as:
- support for [normalize function](https://github.com/ditrit/verdeter/blob/main/docs/normalization/normalization.md), ex: `LowerString` (lower the input string)
- support for [key specific checks](https://github.com/ditrit/verdeter/blob/main/docs/using_it_for_real/using_it_for_real.md), ex: `StringNotEmpty`(check if the input string is empty), `CheckIsHighPort`(check is the input integer is a high tcp port))
- support for constraints, ex: check for specific arch
- support for dynamic default values (named *Computed values*), ex: set `time.Now().Unix()` as a default for a "time" key 


## How Verdeter differ from viper in handling configuration value

Verdeter uses the following precedence order. Each item takes precedence over the item below it:

1. Explicit call to `viper.Set`: 

    `viper.Set(key)` set the key to a fixed value 
    
    *Example: `viper.Set("age", 25)` will set the key "**age**" to `25`*

2. POSIX flags

    Cli flags are handled by cobra using [pflag](https://github.com/spf13/pflag)

    *Example: appending the flag `--age 25` will set the key "**age**" to `25`*

3. Environment variables

    Environment Variable are handled by viper (read more [here](https://github.com/spf13/viper#working-with-environment-variables))

    *Example: running `export <APP_NAME>_age` will export an environment variable (the `<APP_NAME>` is set by verdeter). Verdeter will bind automatically the environment variable name to a viper key when the developer will define the key he needs. Then, when the developer retreive a value for the "**age**" key with a call to `viper.Get("age)`, viper get all the environment variable and find the value of `<APP_NAME>_age`.*
    

4. Value in a config file

    Viper support reading from [JSON, TOML, YAML, HCL, envfile and Java properties config files](https://github.com/spf13/viper#what-is-viper). The developer need to set a key named "**config_path**" to set the path to the config file or the path to the config directory.

    *Example:*
    Let's say the "**config_path**" is set to `./conf.yml` and the file looks like below
    ```yml
    # conf.yml
    author:
        name: bob
    age: 25
    ```
    Then you would use `viper.Get("author.name")` to access the value `bob` and `viper.Get("age")` to access the value `25`.

5. Dynamic default values (*computed values*)

    Verdeter allow the user of "*computed values*" as dynamic default values. It means that the developer can values returned by functions as default values.

    *Example:*
    The function `defaultTime` will provide a unix time integer.

    ```go
    var defaultTime verdeter.models.DefaultValueFunction :=  func () interface{} {
        return time.Now().Unix()
    }
    ```

    We bind this function to the key time using verdeter.

    ```go
    (*VerdeterCommand).SetComputedValue("time", defaultTime)
    ```

    Then the value can be retreived easily using `viper.Get("time")` as usual


6. static default

    Static defaults can be set using verdeter
    ```go
    // of course here the value is static
    (*VerdeterCommand).SetDefault("time", 1661957668)
    ```
    Alternatively you can use viper directly to do exactly the same thing (please note that we will use `(*VerdeterCommand).SetDefault` in the rest of the documentation).
    ```go
    viper.SetDefault("time", 1661957668)
    ```


7. type default (0 for an integer)

    If a key is *not set* and *not marked as required (using `(*VerdeterCommand).SetRequired(<key_name>)`)*, then a call to `viper.Get<Type>(<key_name>)` will return the default value for this `<Type>`.

    *Example:* let's say thay we **did not** call `(*VerdeterCommand).SetRequired("time")` to set the key "time" as required.
    Then a call to  `viper.GetInt("time")` will return `0`. (please note that a call to `viper.Get(<key>)` returns an `interface{}` wich has no "defaut value").


## Basic Example

Let's create a rootCommand named "myApp"
```go

var rootCommand = verdeter.NewConfigCmd(
	// Name of the app 
    "myApp", 
    
    // A short description
    "myApp is an amazing piece of software",
    
    // A longer description
    `myApp is an amazing piece of software,
that everyone can use thanks to verdeter`,

    // Callback
	func(cfg *verdeter.VerdeterCommand, args []string) {
        key := "author.name"
		fmt.Printf("value for %q is %q\n", key, viper.GetString(key))
	})
```

You might to receive args on the command line, set the number of args you want.
If more are provided, Cobra will throw an error.

```go
// only 2 args please
rootCommand.SetNbArgs(2)
``` 

Then I want to add configuration to this command, for example to bind an address and a port to myApp. 

```go
// Adding a local key.
rootCommand.LKey("addr", verdeter.IsStr, "a", "bind to IPV4 addr")
rootCommand.LKey("port", verdeter.IsInt, "p", "bind to TCP port")

/* if you want sub commands to inherit this flag, 
   use (*verdeter.VerdeterCommand).GKey instead */
```

> The config types availables are `verdeter.IsStr`, `verdeter.IsInt`, `verdeter.IsUint` and `verdeter.IsBool`.

A default value can be set for each config key

```go
rootCommand.SetDefault("addr", "127.0.0.1")
rootCommand.SetDefault("port", 7070)
```

A validator can be bound to a config key.

```go
// creating a validator from scratch 
addrValidator := models.Validator{
    // the name of the validator
    Name: "IPV4 validator",

    // the actual validation function
    Func: func (input interface{}) error {
        valueStr, ok := input.(string)
        if !ok {
            return fmt.Error("wrong input type")
        }
        parts := strings.Split(".")
        if len(parts)!=4 {
            return fmt.Errorf("An IPv4 is composed of four 8bit integers, fount  %d", len(parts))
        }
        for _,p := parts {
            intVal, err := strconv.Atoi(p)
            if err != nil {
                return err
            }
            if intVal<0 || intVal >255 {
                return fmt.Error("one of the part in the string is not a byte")
            }
            
        }
    },
}

// using the validator we just created
rootCommand.SetValidator("addr", addrValidator)

// verdeter comes with some predefined validators
rootCommand.SetValidator("port", verdeter.validators.CheckTCPHighPort)
```

Config key can be marked as required. The cobra function [(* cobra.Command).PreRunE](https://pkg.go.dev/github.com/spf13/cobra#Command) will fail if the designated config key is not provided, preventing the callback to run.
```go
rootCommand.SetRequired("addr")
```

To actually run the command, use this code in your main.go

```go
func init(){
     // Initialize the command
    rootCommand.Initialize()
    // setup keys
    // rootCommand.LKey("port", ve......

}
func main() {
   

    /*
        YOUR CODE HERE
    */

    // Launch the command
    rootCommand.Execute()

}
```

## Contributing Guidelines

See [CONTRIBUTING](CONTRIBUTING.md)