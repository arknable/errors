package errors

import (
	"errors"
	"fmt"
)

// Wrap wraps given error
func Wrap(err error) Error {
	werr := new(theError)
	e, ok := err.(Error)
	if !ok {
		werr.code = ErrUnknown
		werr.wrappers = []ErrorScene{}
		werr.message = err.Error()
		werr.scene = getScene()
	} else {
		werr.code = e.Code()
		werr.scene = e.Scene()
		werr.message = e.Message()
		werr.wrappers = append(werr.wrappers, getScene())
	}
	return werr
}

// WrapString wraps given error message
func WrapString(msg string) Error {
	return Wrap(errors.New(msg))
}

// WrapStringf wraps given formatted error message
func WrapStringf(msg string, args ...interface{}) Error {
	return Wrap(fmt.Errorf(msg, args...))
}

// Empty creates error with empty message
func Empty() Error {
	return new(theError)
}
