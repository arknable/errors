package errors

// New creates new Error
func New(msg string) Error {
	err := new(theError)
	err.wrappers = []ErrorScene{}
	err.message = msg
	err.scene = getScene()
	return err
}

// Wrap wraps given error
func Wrap(err error) error {
	e, ok := err.(Error)
	if !ok {
		return err
	}
	werr := new(theError)
	werr.scene = e.Scene()
	werr.message = e.Message()
	werr.wrappers = append(werr.wrappers, getScene())
	return werr
}
