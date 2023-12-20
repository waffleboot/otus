package adapter

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type (
	postgresClient struct {
		conn *pgx.Conn
	}
)

func NewPostgresClient(ctx context.Context, connStr string, connTimeout time.Duration) (*postgresClient, error) {
	if connStr == "" {
		return nil, errors.New("empty conn str")
	}

	timeoutCtx, cancel := context.WithTimeout(ctx, connTimeout)
	defer cancel()

	conn, err := pgx.Connect(timeoutCtx, connStr)
	if err != nil {
		return nil, fmt.Errorf("connect: %w", err)
	}

	err = conn.Ping(ctx)
	if err != nil {
		return nil, fmt.Errorf("ping: %w", err)
	}

	return &postgresClient{conn: conn}, nil
}

func (s *postgresClient) Initialize(ctx context.Context) error {
	_, err := s.conn.Exec(ctx, "create schema if not exists webapp")
	if err != nil {
		return fmt.Errorf("create schema: %w", err)
	}

	_, err = s.conn.Exec(ctx, "create table if not exists webapp.metadata (id uuid)")
	if err != nil {
		return fmt.Errorf("create table: %w", err)
	}
	return nil
}

func (s *postgresClient) Get(ctx context.Context, id uuid.UUID) error {
	err := s.conn.QueryRow(ctx, "select * from webapp.metadata where id=$1", id).Scan(&id)
	if err != nil {
		return fmt.Errorf("query row: %w", err)
	}
	return nil
}

func (s *postgresClient) Put(ctx context.Context, id uuid.UUID) error {
	_, err := s.conn.Exec(ctx, "insert into webapp.metadata (id) values ($1)", id)
	if err != nil {
		return fmt.Errorf("insert: %w", err)
	}
	return nil
}

func (s *postgresClient) All(ctx context.Context) ([]uuid.UUID, error) {
	rows, err := s.conn.Query(ctx, "select * from webapp.metadata")
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

func (s *postgresClient) Close(ctx context.Context) error {
	err := s.conn.Close(ctx)
	if err != nil {
		return fmt.Errorf("close conn: %w", err)
	}

	return nil
}
