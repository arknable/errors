# arknable/errors

Wraps error with function informations where it occured or propagated.

## Install

`go get -u github.com/arknable/errors`

## How To Use

To create a new error use `New()` function as follows,
```
// parsing.go
func parseSomething() error {
    // .......... other thing happen here
    if err := doSomething(); err != nil {
        return errors.New("Missing parameter: foo")
    }    
    // .......... other thing happen here
    return nil
}
```

then if `parseSomething` called within other function, we propagate the error using `Wrap()`:

```
// something.go
function getSomething() error {
    // .......... other thing happen here
    if err := parseSomething(); err != nil {
        return errors.Wrap(err)
    }
    // .......... other thing happen here
}
```

```
fewthings.go
function getFewThings() error {
    // .......... other thing happen here
    if err := getSomething(); err != nil {
        return errors.Wrap(err)
    }
    // .......... other thing happen here
}
```

the result will be something like this,
```
Missing parameter: foo
at /path/to/folder/parsing.go:6 (parseSomething)
at /path/to/folder/something.go:18 (getSomething)
at /path/to/folder/fewthings.go:77 (getFewThings)
```

## License

This project is licensed under the BSD 2 License - see the [LICENSE](LICENSE) file for details.

