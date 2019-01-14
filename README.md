# arknable/errors

Wraps error with function informations where it occured or propagated.

## Install

`go get -u github.com/arknable/errors`

## How To Use

To create a new error, use `New()` as follows,
```go
// parsing.go
func parseSomething() error {
    // .......... other things happen here
    if err := doSomething(); err != nil {
        return errors.New("Missing parameter: foo")
    }    
    // .......... other things happen here
    return nil
}
```

or use `FromError()` to use standard error,
```go
// parsing.go
func parseSomething() error {
    // .......... other things happen here
    if err := doSomething(); err != nil {
        // Lets say returned message is "Missing parameter: foo"
        return errors.FromError(err) 
    }    
    // .......... other things happen here
    return nil
}
```

then if `parseSomething` called within other function, we propagate the error using `Wrap()`:

```go
// something.go
function getSomething() error {
    // .......... other things happen here
    if err := parseSomething(); err != nil {
        return errors.Wrap(err)
    }
    // .......... other things happen here
}
```

```go
fewthings.go
function getFewThings() error {
    // .......... other things happen here
    if err := getSomething(); err != nil {
        return errors.Wrap(err)
    }
    // .......... other things happen here
}
```

the result will be something like this,
```bash
Missing parameter: foo
at /path/to/folder/parsing.go:6 (parseSomething)
at /path/to/folder/something.go:18 (getSomething)
at /path/to/folder/fewthings.go:77 (getFewThings)
```

### JSON

Error can me marshalled on unmarshalled to JSON via standard fashion.

To marshal error,
```go
data, err := json.Marshal(expectedErr)
if err != nil {
    return errors.FromError(err)
}
```

and to unmarshal,
```
resultErr := errs.Empty()
if err := json.Unmarshal(data, resultErr); err != nil {
    return errors.FromError(err)
}
```

the marshalled informations are code and message,
```json
{"code":3,"message":"an error occured"}
```

Other informations ignored because error marshaling mostly used for http response so callers infomation only worth using for logging and debugging.

## License

This project is licensed under the BSD 2 License - see the [LICENSE](LICENSE) file for details.

