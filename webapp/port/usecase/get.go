package usecase

import (
	"context"

	"github.com/google/uuid"
)

type GetFileUseCase interface {
	GetFile(ctx context.Context, id uuid.UUID) ([]byte, error)
}
