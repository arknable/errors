package errors

import "encoding/json"

// Error is a trouble :p.
// It keeps callers, if asked via Wrap().
type Error interface {
	error
	json.Marshaler
	json.Unmarshaler

	// Returns location of the error
	Scene() ErrorScene

	// Returns code of the error
	Code() uint16

	// Sets error code
	WithCode(uint16) Error

	// Returns message of the error
	Message() string

	// Returns functions that propagate the error
	Wrappers() []ErrorScene

	// Returns true if the error wrapped
	HasWrappers() bool
}
