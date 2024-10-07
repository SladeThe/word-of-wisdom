package repositories

import (
	"errors"
)

// A repository method must return either sole repository error or a multi-error consisting of
// exactly one repository error and one or more other errors.

var (
	ErrNotFound = newError("not found")
	ErrInternal = newError("internal error")
)

type Error struct {
	error
}

func newError(msg string) Error {
	return Error{error: errors.New(msg)}
}

func IsError(err error) bool {
	var repositoryErr Error
	return errors.As(err, &repositoryErr)
}
