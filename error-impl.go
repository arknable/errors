package errors

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// Implements Error
type theError struct {
	code     uint16
	scene    ErrorScene
	message  string
	wrappers []ErrorScene
}

// Code implements Error.Code
func (e *theError) Code() uint16 {
	return e.code
}

// Scene implements Error.Scene
func (e *theError) Scene() ErrorScene {
	return e.scene
}

// Message implements Error.Message
func (e *theError) Message() string {
	return e.message
}

// Wrappers implements Error.Wrappers
func (e *theError) Wrappers() []ErrorScene {
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
func (e *theError) HasWrappers() bool {
	return (e.wrappers != nil) && (len(e.wrappers) > 0)
}

// WithCode implements Error.WithCode
func (e *theError) WithCode(code uint16) Error {
	e.code = code
	return e
}

// MarshalJSON implements json.Marshaler
func (e *theError) MarshalJSON() ([]byte, error) {
	err := new(jsError)
	err.Code = e.Code()
	err.Message = e.Message()
	return json.Marshal(err)
}

// UnmarshalJSON implements json.Unmarshaler
func (e *theError) UnmarshalJSON(data []byte) error {
	jerr := new(jsError)
	if err := json.Unmarshal(data, jerr); err != nil {
		return err
	}
	e.code = jerr.Code
	e.message = jerr.Message
	return nil
}

// Error implements error
func (e theError) Error() string {
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
