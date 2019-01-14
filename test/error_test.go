package test

import (
	"encoding/json"
	"errors"
	"testing"

	errs "github.com/arknable/errors"
	"github.com/stretchr/testify/assert"
)

func TestNewError(t *testing.T) {
	err := thirdWrapFunc()
	e, ok := err.(errs.Error)
	assert.True(t, ok)
	assert.NotNil(t, e)
	t.Log(err)
}

func TestFromError(t *testing.T) {
	msg := "standard error"
	e := errors.New(msg)
	err := errs.FromError(e)
	castedErr, ok := err.(errs.Error)
	assert.True(t, ok)
	assert.NotNil(t, castedErr)
	assert.Equal(t, e.Error(), err.Message())
	assert.Equal(t, e.Error(), castedErr.Message())
	assert.Equal(t, err.Message(), castedErr.Message())
}

func TestErrorMessage(t *testing.T) {
	msgErr := "an error occured"
	err := errs.New(msgErr)
	e, ok := err.(errs.Error)
	assert.True(t, ok)
	assert.Equal(t, msgErr, e.Message())
}

func TestErrorCode(t *testing.T) {
	code := uint16(97)
	err := errs.New("an error occured")
	assert.Equal(t, errs.ErrUnknown, err.Code())
	err.WithCode(code)
	assert.Equal(t, code, err.Code())
}

func TestMarshalError(t *testing.T) {
	msg := "an error occured"
	expectedErr := errs.New(msg)
	data, err := json.Marshal(expectedErr)
	if err != nil {
		t.Fatal(err)
	}
	resultErr := errs.New("")
	if err := json.Unmarshal(data, resultErr); err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, expectedErr.Code(), resultErr.Code())
	assert.Equal(t, expectedErr.Message(), resultErr.Message())
}
