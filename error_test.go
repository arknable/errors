package errors

/*
	BSD 2-Clause License

	Copyright (c) 2019, Lufty Abdillah
	All rights reserved.

	Redistribution and use in source and binary forms, with or without
	modification, are permitted provided that the following conditions are met:

	1. Redistributions of source code must retain the above copyright notice, this
	list of conditions and the following disclaimer.

	2. Redistributions in binary form must reproduce the above copyright notice,
	this list of conditions and the following disclaimer in the documentation
	and/or other materials provided with the distribution.

	THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
	AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
	IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
	DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
	FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
	DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
	SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
	CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
	OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
	OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
*/

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

const testErrorMessage = "something is broken"

func errorFunc() error {
	return WrapString(testErrorMessage)
}

func firstWrapFunc() error {
	return Wrap(errorFunc())
}

func secondWrapFunc() error {
	return Wrap(firstWrapFunc())
}

func thirdWrapFunc() error {
	return Wrap(secondWrapFunc())
}

func TestWrapError(t *testing.T) {
	err := thirdWrapFunc()
	e, ok := err.(Error)
	assert.True(t, ok)
	assert.NotNil(t, e)
	assert.Equal(t, e.Message(), testErrorMessage)
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
	code := uint16(97)
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
