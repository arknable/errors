package errors

import (
	"errors"
	"fmt"
)

// ErrUnknown defines an unknown error
const ErrUnknown = -1

// Wrap wraps given error
func Wrap(err error) Error {
	e, ok := err.(Error)
	if !ok {
		ie := new(implError)
		ie.origin = err
		ie.code = ErrUnknown
		ie.wrappers = []ErrorScene{}
		ie.message = err.Error()
		ie.scene = getScene()
		return ie
	}

	if e.Scene() == nil {
		e.setScene(getScene())
	} else {
		e.appendWrapper(getScene())
	}
	return e
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
	err.origin = errors.New(message)
	return err
}

// Newf creates new error with given formatted message
func Newf(format string, v ...interface{}) Error {
	return New(fmt.Sprintf(format, v...))
}

// Is returns true if 'wrapped' is originally 'orig'.
func Is(wrapped, orig error) bool {
	e, ok := wrapped.(Error)
	if !ok {
		return false
	}
	return e.Equal(orig)
}

// Origin returns original error, if wrapped. Otherwise just return err itself.
func Origin(err error) error {
	wrapped, ok := err.(Error)
	if ok {
		return wrapped.Origin()
	}
	return err
}
