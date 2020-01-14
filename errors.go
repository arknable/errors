package errors

import (
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
	return Wrap(New(msg))
}

// WrapStringf wraps given formatted error message
func WrapStringf(msg string, args ...interface{}) Error {
	return WrapString(fmt.Sprintf(msg, args...))
}

// Empty creates error with empty message
func Empty() Error {
	return new(theError)
}

// New creates new error with given message
func New(message string) Error {
	err := new(theError)
	err.code = ErrUnknown
	err.message = message
	return err
}
