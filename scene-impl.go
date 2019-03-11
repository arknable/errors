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
	"fmt"
	"path"
	"runtime"
	"strings"
)

// Implements ErrorScene
type errorScene struct {
	fileName   string
	lineNumber int
	funcName   string
}

// FileName implements ErrorScene.FileName
func (s *errorScene) FileName() string {
	return s.fileName
}

//LineNumber implements ErrorScene.LineNumber
func (s *errorScene) LineNumber() int {
	return s.lineNumber
}

// FuncName implements ErrorScene.FuncName
func (s *errorScene) FuncName() string {
	return s.funcName
}

// Formats scene to string
func sceneToString(s ErrorScene) string {
	return fmt.Sprintf("at %s:%d (%s)\n", s.FileName(), s.LineNumber(), s.FuncName())
}

// returns location of the caller
func getScene() *errorScene {
	pc, fname, line, ok := runtime.Caller(2)
	if !ok {
		return nil
	}
	loc := new(errorScene)
	fn := runtime.FuncForPC(pc)
	if fn != nil {
		loc.funcName = path.Base(fn.Name())
		if strings.Contains(loc.funcName, ".") {
			splitted := strings.Split(loc.funcName, ".")
			loc.funcName = splitted[len(splitted)-1]
		}
	}
	loc.fileName = fname
	loc.lineNumber = line
	return loc
}
