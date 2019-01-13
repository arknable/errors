package errors

// ErrorScene is where an error occured
type ErrorScene interface {
	FileName() string // Full path to the file where error occured
	LineNumber() int  // Line number inside the error file
	FuncName() string // Function name that cause the error
}
