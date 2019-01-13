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
	t.Log(err)
}
