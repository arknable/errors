# arknable/gerror

Wraps error with function informations where it occured or propagated.

## Install

`go get -u github.com/arknable/errors`

## How To Use
Full documentation is on [GoDoc](https://godoc.org/github.com/arknable/errors). To create a new error, use `New()` as follows,
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

Encoding or decoding can be done via standard fashion.

To marshal:
```go
e := errors.New("an error occured").WithCode(3)
data, err := json.Marshal(e)
if err != nil {
    return errors.FromError(err)
}
```

and to unmarshal:
```
e := errors.Empty()
if err := json.Unmarshal(data, e); err != nil {
    return errors.FromError(err)
}
```

the marshaled informations are code and message,
```json
{"code":3,"message":"an error occured"}
```

Other informations ignored because error marshaling mostly used for http response so callers infomation only worth using for logging and debugging .... I think :D.

## License

This project is licensed under the BSD 2 License - see the [LICENSE](LICENSE) file for details.

