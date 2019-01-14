package test

import (
	"encoding/json"
	"testing"

	"github.com/arknable/errors"
	"github.com/stretchr/testify/assert"
)

func TestNewError(t *testing.T) {
	err := thirdWrapFunc()
	e, ok := err.(errors.Error)
	assert.True(t, ok)
	assert.NotNil(t, e)
	t.Log(err)
}

func TestErrorMessage(t *testing.T) {
	msgErr := "an error occured"
	err := errors.New(msgErr)
	e, ok := err.(errors.Error)
	assert.True(t, ok)
	assert.Equal(t, msgErr, e.Message())
}

func TestErrorCode(t *testing.T) {
	code := uint16(97)
	err := errors.New("an error occured")
	assert.Equal(t, errors.ErrUnknown, err.Code())
	err.WithCode(code)
	assert.Equal(t, code, err.Code())
}

func TestMarshalError(t *testing.T) {
	msg := "an error occured"
	expectedErr := errors.New(msg)
	data, err := json.Marshal(expectedErr)
	if err != nil {
		t.Fatal(err)
	}
	resultErr := errors.New("")
	if err := json.Unmarshal(data, resultErr); err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, expectedErr.Code(), resultErr.Code())
	assert.Equal(t, expectedErr.Message(), resultErr.Message())
}
