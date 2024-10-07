package services

import (
	"context"
	"errors"
)

type contextKey string

var (
	ErrMissingServices = errors.New("missing services")

	servicesKey = contextKey("wow.services")
)

type Services struct {
	Challenge    Challenge
	WordOfWisdom WordOfWisdom
}

func New(
	challenge Challenge,
	wordOfWisdom WordOfWisdom,
) Services {
	return Services{
		Challenge:    challenge,
		WordOfWisdom: wordOfWisdom,
	}
}

func Set(ctx context.Context, ss Services) context.Context {
	return context.WithValue(ctx, servicesKey, ss)
}

func Get(ctx context.Context) (Services, error) {
	ss, ok := ctx.Value(servicesKey).(Services)
	if !ok {
		return Services{}, ErrMissingServices
	}
	return ss, nil
}

func Must(ctx context.Context) Services {
	ss, err := Get(ctx)
	if err != nil {
		panic(err)
	}
	return ss
}
