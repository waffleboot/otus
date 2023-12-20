package usecase

import (
	"context"

	"github.com/google/uuid"
)

type CreateFileUseCase interface {
	CreateFile(ctx context.Context, content []byte) (uuid.UUID, error)
}
