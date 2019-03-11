# arknable/errors

Wraps error with function informations where it occured or propagated.

## How To Use
Full documentation is on [GoDoc](https://godoc.org/github.com/arknable/errors). Use `Wrap` to wrap an `error` object or either `WrapString` or `WrapStringf` to create a new `error`.
```go
// parsing.go
func parseSomething() error {
    // .......... other things happen here
    if err := doSomething(); err != nil {
        return errors.WrapString("Missing parameter: foo")
    }    
    // .......... other things happen here
    return nil
}
```

or for `error` object,
```go
// parsing.go
func parseSomething() error {
    // .......... other things happen here
    if err := doSomething(); err != nil {
        return errors.Wrap(err) 
    }    
    // .......... other things happen here
    return nil
}
```

assume `parseSomething` called from other function:

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
// fewthings.go
function getFewThings() error {
    // .......... other things happen here
    if err := getSomething(); err != nil {
        return errors.Wrap(err)
    }
    // .......... other things happen here
}
```

the result will be something like this:
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
e := errors.WrapString("an error occured").WithCode(3)
data, err := json.Marshal(e)
if err != nil {
    return errors.Wrap(err)
}
```

and to unmarshal:
```
e := errors.Empty()
if err := json.Unmarshal(data, e); err != nil {
    return errors.Wrap(err)
}
```

the marshaled informations are code and message,
```json
{"code":3,"message":"an error occured"}
```

## License

This project is licensed under the BSD 2 License - see the [LICENSE](LICENSE) file for details.

