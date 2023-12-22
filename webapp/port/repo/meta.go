package repo

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type Metadata interface {
	Get(ctx context.Context, id uuid.UUID) error
	Put(ctx context.Context, id uuid.UUID, at time.Time) error
	All(ctx context.Context) ([]uuid.UUID, error)
}
