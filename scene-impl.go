package errors

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
