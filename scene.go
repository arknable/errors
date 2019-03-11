package errors

// ErrorScene is where an error occured
type ErrorScene interface {
	// Returns full path to the file where error occured
	FileName() string

	// Returns line number inside the error file
	LineNumber() int

	// Returns function name that cause the error
	FuncName() string
}
