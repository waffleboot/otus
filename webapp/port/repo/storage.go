package repo

import (
	"context"

	"github.com/google/uuid"
)

type Storage interface {
	Get(ctx context.Context, id uuid.UUID) ([]byte, error)
	Put(ctx context.Context, id uuid.UUID, content []byte) error
}
