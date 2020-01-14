package errors

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// Implements Error
type implError struct {
	code     int
	scene    ErrorScene
	message  string
	wrappers []ErrorScene
}

// Code implements Error.Code
func (e *implError) Code() int {
	return e.code
}

// Scene implements Error.Scene
func (e *implError) Scene() ErrorScene {
	return e.scene
}

// Message implements Error.Message
func (e *implError) Message() string {
	return e.message
}

// Wrappers implements Error.Wrappers
func (e *implError) Wrappers() []ErrorScene {
	length := len(e.wrappers)
	if length != 0 {
		return nil
	}
	wrappers := make([]ErrorScene, length)
	for _, v := range e.wrappers {
		wrappers = append(wrappers, v)
	}
	return wrappers
}

// HasWrappers implements Error.HasWrappers
func (e *implError) HasWrappers() bool {
	return (e.wrappers != nil) && (len(e.wrappers) > 0)
}

// WithCode implements Error.WithCode
func (e *implError) WithCode(code int) Error {
	e.code = code
	return e
}

// MarshalJSON implements json.Marshaler
func (e *implError) MarshalJSON() ([]byte, error) {
	err := new(jsError)
	err.Code = e.Code()
	err.Message = e.Message()
	return json.Marshal(err)
}

// UnmarshalJSON implements json.Unmarshaler
func (e *implError) UnmarshalJSON(data []byte) error {
	jerr := new(jsError)
	if err := json.Unmarshal(data, jerr); err != nil {
		return err
	}
	e.code = jerr.Code
	e.message = jerr.Message
	return nil
}

// Error implements error
func (e implError) Error() string {
	if e.scene == nil {
		return e.message
	}
	str := bytes.Buffer{}
	str.WriteString(fmt.Sprintf("%s\n", e.message))
	str.WriteString(sceneToString(e.scene))

	if len(e.wrappers) > 0 {
		for _, c := range e.wrappers {
			str.WriteString(sceneToString(c))
		}
	}
	return str.String()
}

// Equal implements Error.Equal
func (e implError) Equal(err error) bool {
	return e.message == err.Error()
}
