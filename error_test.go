package errors

import (
	"encoding/json"
	goerr "errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var tErrorMessage = goerr.New("something is broken")

func tErrorFunc() error {
	return WrapString(tErrorMessage.Error())
}

func tFirstWrapFunc() error {
	return Wrap(tErrorFunc())
}

func tSecondWrapFunc() error {
	return Wrap(tFirstWrapFunc())
}

func tThirdWrapFunc() error {
	return Wrap(tSecondWrapFunc())
}

func TestWrapError(t *testing.T) {
	err := tThirdWrapFunc()
	e, ok := err.(Error)
	assert.True(t, ok)
	assert.NotNil(t, e)
	assert.Equal(t, e.Message(), tErrorMessage.Error())
	assert.True(t, e.Equal(tErrorMessage))
}

func TestWrapString(t *testing.T) {
	msg := "standard error"
	err := WrapString(msg)
	assert.NotNil(t, err)
	assert.Equal(t, msg, err.Message())
}

func TestWrapFormattedString(t *testing.T) {
	msg := "standard error from %s"
	err := WrapStringf(msg, "Google")
	assert.NotNil(t, err)
	assert.Equal(t, fmt.Sprintf(msg, "Google"), err.Message())
}

func TestErrorCode(t *testing.T) {
	code := 97
	err := WrapString("an error occured")
	assert.Equal(t, ErrUnknown, err.Code())
	err.WithCode(code)
	assert.Equal(t, code, err.Code())
}

func TestMarshalError(t *testing.T) {
	msg := "an error occured"
	expectedErr := WrapString(msg).WithCode(3)
	data, err := json.Marshal(expectedErr)
	if err != nil {
		t.Fatal(err)
	}
	resultErr := Empty()
	if err := json.Unmarshal(data, resultErr); err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, expectedErr.Code(), resultErr.Code())
	assert.Equal(t, expectedErr.Message(), resultErr.Message())
}
