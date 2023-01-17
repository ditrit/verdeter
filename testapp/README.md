# TestApp for verdeter

This golang module is a test app for verdeter.
This test app use all the major feature of verdeter.

This app and this file will be kept up to date with verdeter development.
You can see some bash lines and their associated outputs.

```txt
$ MATH_ORGANISATION_NAME=ditrit go run main.go
The root command does nothing but print the config key 'organisation.name'
value for "key=organisation.name": ditrit
```

```txt
$ go run main.go
The root command does nothing but print the config key 'organisation.name'
value for "key=organisation.name": orness
```

```txt
$ go run main.go add --int1 -1
Some errors were collected during initialization:
- validator "Superior to 0" failed for key "int1": "under1"
- "int2" is required
Error: validation failed (2 errors)
Usage:
  math add [flags]

Flags:
  -h, --help       help for add
      --int1 int   Integer 1
      --int2 int   Integer 2

Global Flags:
      --config_path string         Path to the config directory/file
      --organisation.name string   An organisation name

panic: validation failed (2 errors)

...
```

```txt
$ go run main.go add --int1 12 --int2 5
print the result of --int1 + --int2
result: 17
```