package test

import "github.com/arknable/errors"

const errorMessage = "something is broken"

func errorFunc() error {
	return errors.WrapString(errorMessage)
}

func firstWrapFunc() error {
	return errors.Wrap(errorFunc())
}

func secondWrapFunc() error {
	return errors.Wrap(firstWrapFunc())
}

func thirdWrapFunc() error {
	return errors.Wrap(secondWrapFunc())
}
