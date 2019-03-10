package test

import "github.com/arknable/gerror"

func errorFunc() error {
	return gerror.New("something is broken")
}

func firstWrapFunc() error {
	return errorFunc()
}

func secondWrapFunc() error {
	return gerror.Wrap(firstWrapFunc())
}

func thirdWrapFunc() error {
	return gerror.Wrap(secondWrapFunc())
}
