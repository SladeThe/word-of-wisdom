package services

import (
	"errors"
)

// A service method must return either sole service error or a multi-error consisting
// of exactly one service error and one or more other errors.

var (
	ErrInvalidArguments  = newError("invalid arguments")
	ErrNotFound          = newError("not found")
	ErrServiceOverloaded = newError("service overloaded")
	ErrInternal          = newError("internal error")
)

type Error struct {
	error
}

func newError(msg string) Error {
	return Error{error: errors.New(msg)}
}

func IsError(err error) bool {
	var serviceErr Error
	return errors.As(err, &serviceErr)
}
