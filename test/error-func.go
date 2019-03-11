package test

import "github.com/arknable/gerror"

const errorMessage = "something is broken"

func errorFunc() error {
	return gerror.WrapString(errorMessage)
}

func firstWrapFunc() error {
	return gerror.Wrap(errorFunc())
}

func secondWrapFunc() error {
	return gerror.Wrap(firstWrapFunc())
}

func thirdWrapFunc() error {
	return gerror.Wrap(secondWrapFunc())
}
