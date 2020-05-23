package errors

import (
	"fmt"
)

// ErrUnknown defines an unknown error
const ErrUnknown = -1

// Wrap wraps given error
func Wrap(err error) Error {
	werr := new(implError)
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
	return new(implError)
}

// New creates new error with given message
func New(message string) Error {
	err := new(implError)
	err.code = ErrUnknown
	err.message = message
	return err
}

// Newf creates new error with given formatted message
func Newf(format string, v ...interface{}) Error {
	return New(fmt.Sprintf(format, v...))
}
