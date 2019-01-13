package errors

// Error is a trouble :p.
// It keeps callers, if asked via Wrap().
type Error interface {
	error

	Scene() ErrorScene      // Location of the error
	Message() string        // Message of the error
	Wrappers() []ErrorScene // Functions that propagate the error
	HasWrappers() bool      // Convenient way to check whether the error wrapped
}
