package repositories

import (
	"context"
	"errors"
)

type contextKey string

var (
	ErrMissingRepositories = errors.New("missing repositories")

	repositoriesKey = contextKey("wow.repositories")
)

type Repositories struct {
	Client       Client
	WordOfWisdom WordOfWisdom
}

func New(
	client Client,
	wordOfWisdom WordOfWisdom,
) Repositories {
	return Repositories{
		Client:       client,
		WordOfWisdom: wordOfWisdom,
	}
}

func Set(ctx context.Context, rr Repositories) context.Context {
	return context.WithValue(ctx, repositoriesKey, rr)
}

func Get(ctx context.Context) (Repositories, error) {
	rr, ok := ctx.Value(repositoriesKey).(Repositories)
	if !ok {
		return Repositories{}, ErrMissingRepositories
	}
	return rr, nil
}

func Must(ctx context.Context) Repositories {
	rr, err := Get(ctx)
	if err != nil {
		panic(err)
	}
	return rr
}
