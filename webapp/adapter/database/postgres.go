package database

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

type (
	postgresClient struct {
		pool *pgxpool.Pool
	}
)

func NewPostgresClient(ctx context.Context, connStr string, connTimeout time.Duration) (*postgresClient, error) {
	if connStr == "" {
		return nil, errors.New("empty conn str")
	}

	config, err := pgxpool.ParseConfig(connStr)
	if err != nil {
		return nil, fmt.Errorf("parse config: %w", err)
	}

	timeoutCtx, cancel := context.WithTimeout(ctx, connTimeout)
	defer cancel()

	pool, err := pgxpool.ConnectConfig(timeoutCtx, config)
	if err != nil {
		return nil, fmt.Errorf("connect: %w", err)
	}

	err = pool.Ping(ctx)
	if err != nil {
		return nil, fmt.Errorf("ping: %w", err)
	}

	return &postgresClient{pool: pool}, nil
}

func (s *postgresClient) Initialize(ctx context.Context) error {
	_, err := s.pool.Exec(ctx, "create table if not exists webapp.metadata (id uuid, created_at timestamptz)")
	if err != nil {
		return fmt.Errorf("create table: %w", err)
	}
	_, err = s.pool.Exec(ctx, "delete from webapp.metadata")
	if err != nil {
		return fmt.Errorf("delete all: %w", err)
	}
	return nil
}

func (s *postgresClient) Get(ctx context.Context, id uuid.UUID) error {
	err := s.pool.QueryRow(ctx, "select id from webapp.metadata where id=$1", id).Scan(&id)
	if err != nil {
		return fmt.Errorf("query row: %w", err)
	}
	return nil
}

func (s *postgresClient) Put(ctx context.Context, id uuid.UUID, at time.Time) error {
	_, err := s.pool.Exec(ctx, "insert into webapp.metadata (id,created_at) values ($1,$2)", id, at)
	if err != nil {
		return fmt.Errorf("insert: %w", err)
	}
	return nil
}

func (s *postgresClient) All(ctx context.Context) ([]uuid.UUID, error) {
	rows, err := s.pool.Query(ctx, "select id from webapp.metadata order by created_at desc")
	if err != nil {
		return nil, fmt.Errorf("query: %w", err)
	}
	defer rows.Close()

	var ids []uuid.UUID

	for rows.Next() {
		var id uuid.UUID
		err := rows.Scan(&id)
		if err != nil {
			return nil, fmt.Errorf("scan: %w", err)
		}
		ids = append(ids, id)
	}
	err = rows.Err()
	if err != nil {
		return nil, fmt.Errorf("scan: %w", err)
	}

	return ids, nil
}

func (s *postgresClient) Close() {
	s.pool.Close()
}
