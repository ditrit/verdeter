# Using offset for distributed systems

Let's say you are developping a distributed app. You want your nodes to run on different ports.

Lets start from a basic VerdeterCommand
```go
var verdeterRootCmd = verdeter.NewVerdeterCommand(
	"verdeterapp",
	"verdeterapp is a dummy CLI app",
	"/* Insert a longer description here*/",
	func(cfg *verdeter.VerdeterCommand, args []string) error {
		fmt.Println(verdeter.GetInstanceName())
		// no error to return
		return nil
})
```

Declare the "offset" key.

```go
	verdeterRootCmd.GKey("offset", verdeter.IsInt, "o", "the offset")
```

Now you are ready to use the offset system.
Make sure that you use a different offset for each node of your app. Verdeter do not come with protection on that side.

## Now that the offset system is working, what does it changes ? 

- The prefix of the environment variable changes from `<APP_NAME>_` to `<APP_NAME>_<OFFSET>_` 
- The default name of the config file change from  `<APP_NAME>.ext` to `<APP_NAME>_<OFFSET>.ext` 

This system allow the use of several nodes of your app on the same system.