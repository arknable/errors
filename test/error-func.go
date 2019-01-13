package test

import "github.com/arknable/errors"

func errorFunc() error {
	return errors.New("something is broken")
}

func firstWrapFunc() error {
	return errorFunc()
}

func secondWrapFunc() error {
	return errors.Wrap(firstWrapFunc())
}

func thirdWrapFunc() error {
	return errors.Wrap(secondWrapFunc())
}
