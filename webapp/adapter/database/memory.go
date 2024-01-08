package database

import (
	"context"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"time"

	"github.com/google/uuid"
)

type memoryDatabase struct {
	dir string
}

func NewMemoryDatabase(dir string) (*memoryDatabase, error) {
	d := filepath.Join(dir, "metadata")
	err := os.MkdirAll(d, 0700)
	if err != nil {
		return nil, fmt.Errorf("mkdir: %w", err)
	}
	return &memoryDatabase{dir: d}, nil
}

func (s *memoryDatabase) Get(ctx context.Context, id uuid.UUID) error {
	n := filepath.Join(s.dir, id.String())
	_, err := os.Stat(n)
	if err != nil {
		return fmt.Errorf("file stat: %w", err)
	}
	return nil
}

func (s *memoryDatabase) Put(ctx context.Context, id uuid.UUID, at time.Time) (err error) {
	n := filepath.Join(s.dir, id.String())
	f, err := os.Create(n)
	if err != nil {
		return fmt.Errorf("create file: %w", err)
	}
	defer func() {
		errClose := f.Close()
		if errClose == nil {
			return
		}
		errClose = fmt.Errorf("close: %w", errClose)
		switch {
		case errClose != nil && err != nil:
			err = errors.Join(err, errClose)
		case errClose != nil:
			err = errClose
		}
	}()

	c := at.Format(time.RFC3339)
	_, err = f.Write([]byte(c))
	if err != nil {
		return fmt.Errorf("write: %w", err)
	}

	return nil
}

func (s *memoryDatabase) All(ctx context.Context) ([]uuid.UUID, error) {
	type item struct {
		id uuid.UUID
		at time.Time
	}
	var items []item
	_ = filepath.Walk(s.dir, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return nil
		}

		n := filepath.Base(path)
		id, err := uuid.Parse(n)
		if err != nil {
			return nil
		}
		f, err := os.Open(n)
		if err != nil {
			return nil
		}
		defer func() {
			_ = f.Close()
		}()

		b, err := io.ReadAll(f)
		if err != nil {
			return nil
		}

		at, err := time.Parse(time.RFC3339, string(b))
		if err != nil {
			return nil
		}

		items = append(items, item{id: id, at: at})

		return nil
	})

	sort.Slice(items, func(i, j int) bool {
		return items[i].at.Before(items[j].at)
	})

	ret := make([]uuid.UUID, 0, len(items))
	for _, it := range items {
		ret = append(ret, it.id)
	}

	return ret, nil
}
