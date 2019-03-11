package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/arknable/errors"
	"github.com/stretchr/testify/assert"
)

func TestWrapError(t *testing.T) {
	err := thirdWrapFunc()
	e, ok := err.(errors.Error)
	assert.True(t, ok)
	assert.NotNil(t, e)
	assert.Equal(t, e.Message(), errorMessage)
	t.Log(err)
}

func TestWrapString(t *testing.T) {
	msg := "standard error"
	err := errors.WrapString(msg)
	assert.NotNil(t, err)
	assert.Equal(t, msg, err.Message())
}

func TestWrapFormattedString(t *testing.T) {
	msg := "standard error from %s"
	err := errors.WrapStringf(msg, "Google")
	assert.NotNil(t, err)
	assert.Equal(t, fmt.Sprintf(msg, "Google"), err.Message())
}

func TestErrorCode(t *testing.T) {
	code := uint16(97)
	err := errors.WrapString("an error occured")
	assert.Equal(t, errors.ErrUnknown, err.Code())
	err.WithCode(code)
	assert.Equal(t, code, err.Code())
}

func TestMarshalError(t *testing.T) {
	msg := "an error occured"
	expectedErr := errors.WrapString(msg).WithCode(3)
	data, err := json.Marshal(expectedErr)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(bytes.NewBuffer(data))
	resultErr := errors.Empty()
	if err := json.Unmarshal(data, resultErr); err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, expectedErr.Code(), resultErr.Code())
	assert.Equal(t, expectedErr.Message(), resultErr.Message())
}
