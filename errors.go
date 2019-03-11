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
	"errors"
	"fmt"
)

// Wrap wraps given error
func Wrap(err error) Error {
	werr := new(theError)
	e, ok := err.(Error)
	if !ok {
		werr.code = ErrUnknown
		werr.wrappers = []ErrorScene{}
		werr.message = err.Error()
		werr.scene = getScene()
	} else {
		werr.code = e.Code()
		werr.scene = e.Scene()
		werr.message = e.Message()
		werr.wrappers = append(werr.wrappers, getScene())
	}
	return werr
}

// WrapString wraps given error message
func WrapString(msg string) Error {
	return Wrap(errors.New(msg))
}

// WrapStringf wraps given formatted error message
func WrapStringf(msg string, args ...interface{}) Error {
	return Wrap(fmt.Errorf(msg, args...))
}

// Empty creates error with empty message
func Empty() Error {
	return new(theError)
}
