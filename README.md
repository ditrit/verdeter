# Verdeter

[![GitHub Workflow Status](https://img.shields.io/github/actions/workflow/status/ditrit/verdeter/CI.yml?branch=main&style=flat-square)](https://github.com/spf13/viper/actions?query=workflow%3ACI)
[![Go Report Card](https://goreportcard.com/badge/github.com/ditrit/verdeter?style=flat-square)](https://goreportcard.com/report/github.com/ditrit/verdeter)
![Go Version](https://img.shields.io/badge/go%20version-%3E=1.18-61CFDD.svg?style=flat-square)
[![PkgGoDev](https://pkg.go.dev/badge/mod/github.com/ditrit/verdeter)](https://pkg.go.dev/mod/github.com/ditrit/verdeter)

Verdeter is a library to write CLIs and configuration easily by bringing the power of [cobra](https://github.com/spf13/cobra) and [viper](https://github.com/spf13/viper) in a single library.

Verdeter allow developers to bind a posix compliant flag, an environment variable and a variable in a config file to a viper key with a single line of code.
Verdeter provide a consistent precedence order by extending [viper precedence order](https://github.com/spf13/viper#why-viper).
Verdeter also comes with extra features such as:

- support for [normalize function](#normalization) to normalize user inputs (ex: [`LowerString`] lower the input string)
- support for [key specific validation](#validation), ex: `StringNotEmpty`(check if the input string is empty), `CheckIsHighPort`(check is the input integer is a high tcp port), or `AuthorizedValues`(check if the value of a config key is contained in a defined array of authorized values))
- support for constraints, ex: check for specific arch
- support for dynamic default values (named *Computed values*), ex: set `time.Now().Unix()` as a default for a "time" key.

Table of contents:

- [How Verdeter differ from viper in handling configuration values](#how-verdeter-differ-from-viper-in-handling-configuration-values)
- [Get Started](#get-started)
- [Normalization](#normalization)
- [Validation](#validation)
- [Licence](#licence)
- [Contributing Guidelines](#contributing-guidelines)

## How Verdeter differ from viper in handling configuration values

Verdeter uses the following precedence order. Each item takes precedence over the item below it:

1. Explicit call to `viper.Set`

    `viper.Set(key)` set the key to a fixed value.

    *Example: `viper.Set("age", 25)` will set the key "**age**" to `25`*

2. POSIX flags

    Cli flags are handled by cobra using [pflag](https://github.com/spf13/pflag).

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

    Then the value can be retreived easily using `viper.Get("time")` as usual.

6. static default

    Static defaults can be set using verdeter.

    ```go
    // of course here the value is static
    (*VerdeterCommand).SetDefault("time", 1661957668)
    ```

7. type default (0 for an integer)

    If a key is *not set* and *not marked as required (using `(*VerdeterCommand).SetRequired(<key_name>)`)*, then a call to `viper.Get<Type>(<key_name>)` will return the default value for this `<Type>`.

    *Example:* let's say thay we **did not** call `(*VerdeterCommand).SetRequired("time")` to set the key "time" as required.
    Then a call to  `viper.GetInt("time")` will return `0`. (please note that a call to `viper.Get(<key>)` returns an `interface{}` wich has no "defaut value").

## Get Started

Get the library.

```shell
go get github.com/ditrit/verdeter
```

Let's create an app that print stuff to the terminal. Let's start with the classic "hello world!"

```go

import "github.com/ditrit/verdeter"

var helloCommand = verdeter.BuildVerdeterCommand(verdeter.VerdeterConfig{
	Use:         "hello",
	Long:        "hello is an app that says hello",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Hello World! (from: %q)\n", viper.GetString("from"))
	},
})
```

If you take a look at [verdeter.VerdeterConfig's documentation]() you will observe that it's quite similar to the the original cobra.Command type. That's done on purpose to help you transition easily from your cobra app.

Then I want to add configuration to this command: let's create a config key named "from" that will represent the name of the sender.

```go
// Adding a local key.
helloCommand.GKey("from", verdeter.IsStr, "f", "the sender name")

/* if you want sub commands to inherit this flag/config key, 
   use (*verdeter.VerdeterCommand).GKey instead */
```

> The config types availables are `verdeter.IsStr`, `verdeter.IsInt`, `verdeter.IsUint` and `verdeter.IsBool`.

A default value can be set for that config key. See [.SetDefault() doc]().

```go
// The default value of the config key "from" is nom "earth".
rootCommand.SetDefault("from", "earth")
```

Config key can be marked as required. The cobra function [(* cobra.Command).PreRunE](https://pkg.go.dev/github.com/spf13/cobra#Command) will fail if the designated config key is not provided, preventing your `Run` or `RunE` function to be called.

```go
rootCommand.SetRequired("from")
```

To actually run the command, use the `Execute()` method.

```go
func main() {
   

    /*
        YOUR CODE HERE
    */

    // Launch the command
    helloCommand.Execute()
}
```

## Normalization

Let's say you are building an app that take strings as config values. Instead of asking your user to use only lowercase strings you could set a normalizer with verdeter that will ensure that the string value you will retrieve is actually a lowercase value.

Please note that normalization functions use a specific signature

```go
import "github.com/ditrit/verdeter/models"

var LowerString models.NormalizationFunction = func(val interface{}) interface{} {
	strVal, ok := val.(string)
	if !ok {
		return val
	}
	return strings.ToLower(strVal)
}

verdeterCommand.SetNormalize("keyname", LowerString)
```

---

*The `LowerString` normalization function is actually available at `verdeter.normalization.LowerString`*

## Validation

Let's say you are building an app that serve content over http. You will need to bind your app to a port on the server. You will likely put that in a config key. In order to prevent configuration mistakes that would prevent the application from running, you want to make sure the port number is a TCP high port. Note that Verdeter introduce a validation step that run before the `PreRun` or `PreRunE` function, depending on wich you are using.

First write your validator using verdeter model.

```go
import "github.com/ditrit/verdeter/models"

var validatorTCPHighPort = models.Validator{

    // Give your validator a name
    Name: "TCP High Port Check"

    // Then provide the validation function.
    // (Please note that you need that exact signature)
    Func: func (input interface{}) error {
        // First make sure this is an integer
        portNumber, ok := input.(int)
        if !ok { // seems that is not an integer
            return fmt.Errorf("should be an integer")
        }

        // Check if the port is in the correct interval
        if intVal >= 1024 && intVal <= 65535 {
            return nil
        }
        return fmt.Errorf("value (%d) is not a TCP high port ", port)
    }
} 

// Then register the validator for the config key
verdeterCommand.AddValidator("port", validatorTCPHighPort)
```

---

*The `CheckTCPHighPort` validator is actually available in verdeter at `verdeter.validators.CheckTCPHighPort`*


## Licence

Verdeter is licenced under the Mozilla Public License Version 2.0: see [LICENSE](LICENSE).

## Contributing Guidelines

See [CONTRIBUTING](CONTRIBUTING.md).
