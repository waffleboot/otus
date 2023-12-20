package usecase

import (
	"context"

	"github.com/google/uuid"
)

type ListFilesUseCase interface {
	GetFiles(ctx context.Context) ([]uuid.UUID, error)
}
