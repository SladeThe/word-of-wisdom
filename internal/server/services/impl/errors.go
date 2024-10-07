package impl

import (
	"errors"
	"log"

	"github.com/SladeThe/word-of-wisdom/internal/server/repositories"
	"github.com/SladeThe/word-of-wisdom/internal/server/services"
)

func handleError(err error, msg string) error {
	if err == nil || services.IsError(err) {
		return err
	}

	if errors.Is(err, repositories.ErrNotFound) {
		return services.ErrNotFound
	}

	if err == repositories.ErrInternal {
		log.Println("[ERROR]", msg)
	} else {
		log.Printf("[ERROR] %v: %v", msg, err)
	}

	return services.ErrInternal
}
