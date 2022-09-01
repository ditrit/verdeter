# Normalization

Verdeter support normalization functions.

Let's say you are building an app that take strings as config values. Instead of asking you user to use only lowercase strings you could set a normalizer with verdeter that will ensure that the string value you will retrieve is actually a lowercase value. 

```go
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