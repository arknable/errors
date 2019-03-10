package gerror

// New creates new Error
func New(msg string) Error {
	err := new(theError)
	err.code = ErrUnknown
	err.wrappers = []ErrorScene{}
	err.message = msg
	err.scene = getScene()
	return err
}

// FromError creates new Error from standard error
func FromError(err error) Error {
	return New(err.Error())
}

// Empty creates error with no message
func Empty() Error {
	return New("")
}

// Wrap wraps given error
func Wrap(err error) error {
	e, ok := err.(Error)
	if !ok {
		return err
	}
	werr := new(theError)
	werr.code = e.Code()
	werr.scene = e.Scene()
	werr.message = e.Message()
	werr.wrappers = append(werr.wrappers, getScene())
	return werr
}
