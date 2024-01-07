package database

import (
	"context"
	"errors"
	"sync"
	"time"

	"github.com/google/uuid"
)

type memoryDatabase struct {
	mu  *sync.RWMutex
	dat []uuid.UUID
	idx map[uuid.UUID]struct{}
}

func NewMemoryDatabase() *memoryDatabase {
	return &memoryDatabase{mu: new(sync.RWMutex), idx: make(map[uuid.UUID]struct{})}
}

func (s *memoryDatabase) Get(ctx context.Context, id uuid.UUID) error {
	s.mu.RLock()
	defer s.mu.RUnlock()

	_, ok := s.idx[id]
	if !ok {
		return errors.New("not found")
	}

	return nil
}

func (s *memoryDatabase) Put(ctx context.Context, id uuid.UUID, at time.Time) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.idx[id] = struct{}{}
	s.dat = append(s.dat, id)
	return nil
}

func (s *memoryDatabase) All(ctx context.Context) ([]uuid.UUID, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.dat[:len(s.dat):len(s.dat)], nil
}
