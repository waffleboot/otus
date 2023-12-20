package files

import (
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/google/uuid"
)

type Storage struct {
	staticDir string
}

func NewStorage(staticDir string) (*Storage, error) {
	if staticDir == "" {
		return nil, errors.New("empty static dir config")
	}

	info, err := os.Stat(staticDir)
	if err != nil {
		return nil, fmt.Errorf("static dir info: %w", err)
	}
	if !info.IsDir() {
		return nil, fmt.Errorf("%s is not directory", staticDir)
	}

	return &Storage{staticDir: staticDir}, nil
}

func (s *Storage) Get(ctx context.Context, id uuid.UUID) ([]byte, error) {
	f, err := os.Open(s.filename(id))
	if err != nil {
		return nil, fmt.Errorf("open file: %w", err)
	}
	defer func() {
		errClose := f.Close()
		switch {
		case err != nil && errClose != nil:
			err = errors.Join(err, errClose)
		case errClose != nil:
			err = errClose
		}
	}()

	data, err := io.ReadAll(f)
	if err != nil {
		return nil, fmt.Errorf("write: %w", err)
	}

	return data, nil
}

func (s *Storage) Put(ctx context.Context, id uuid.UUID, content []byte) (err error) {
	f, err := os.Create(s.filename(id))
	if err != nil {
		return fmt.Errorf("create file: %w", err)
	}
	defer func() {
		errClose := f.Close()
		switch {
		case err != nil && errClose != nil:
			err = errors.Join(err, errClose)
		case errClose != nil:
			err = errClose
		}
	}()

	_, err = f.Write(content)
	if err != nil {
		return fmt.Errorf("write: %w", err)
	}

	return nil
}

func (s *Storage) filename(id uuid.UUID) string {
	return filepath.Join(s.staticDir, fmt.Sprintf("%s.dat", id))
}
