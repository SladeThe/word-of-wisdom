package repositories

import (
	"context"

	"github.com/SladeThe/word-of-wisdom/internal/common/entities"
)

type Client interface {
	OneByID(ctx context.Context, id entities.ClientID) (entities.Client, error)
}
