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
	Code() int

	// Sets error code
	WithCode(int) Error

	// Returns message of the error
	Message() string

	// Returns functions that propagate the error
	Wrappers() []ErrorScene

	// Returns true if the error wrapped
	HasWrappers() bool

	// Equal checks if given error equal original error
	Equal(error) bool
}
