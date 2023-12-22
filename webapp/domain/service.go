package domain

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/waffleboot/app/port/repo"
)

type Service struct {
	metadata repo.Metadata
	storage  repo.Storage
}

func NewService(metadata repo.Metadata, storage repo.Storage) *Service {
	return &Service{
		metadata: metadata,
		storage:  storage,
	}
}

func (s *Service) CreateFile(ctx context.Context, content []byte) (uuid.UUID, error) {
	id := uuid.New()

	err := s.storage.Put(ctx, id, content)
	if err != nil {
		return uuid.Nil, fmt.Errorf("put file: %w", err)
	}

	err = s.metadata.Put(ctx, id, time.Now())
	if err != nil {
		return uuid.Nil, fmt.Errorf("add metadata: %w", err)
	}

	return id, nil
}

func (s *Service) GetFile(ctx context.Context, id uuid.UUID) ([]byte, error) {
	err := s.metadata.Get(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("metadata: %w", err)
	}

	return s.storage.Get(ctx, id)
}

func (s *Service) GetFiles(ctx context.Context) ([]uuid.UUID, error) {
	return s.metadata.All(ctx)
}
