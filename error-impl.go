package errors

import (
	"bytes"
	"fmt"
)

type theError struct {
	scene    ErrorScene
	message  string
	wrappers []ErrorScene
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
	return len(e.wrappers) > 0
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
