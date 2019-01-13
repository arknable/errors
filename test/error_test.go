package test

import (
	"testing"

	"github.com/arknable/errors"
	"github.com/stretchr/testify/assert"
)

func TestNewError(t *testing.T) {
	err := thirdWrapFunc()
	e, ok := err.(errors.Error)
	assert.True(t, ok)
	assert.True(t, e.HasWrappers())
}

func TestErrorMessage(t *testing.T) {
	msgErr := "an error occured"
	err := errors.New(msgErr)
	e, ok := err.(errors.Error)
	assert.True(t, ok)
	assert.Equal(t, msgErr, e.Message())
}

func TestErrorCode(t *testing.T) {
	var code uint16 = 97
	err := errors.New("an error occured").WithCode(code)
	e, ok := err.(errors.Error)
	assert.True(t, ok)
	assert.Equal(t, code, e.Code())
}
