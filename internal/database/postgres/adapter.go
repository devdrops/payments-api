package postgres

import (
	"context"

	"payments-api/internal/database"
)

type PostgresAdapter struct {
	conn *database.Client
}

func NewPostgresAdapter(c *database.Client) *Adapter {
	return &PostgresAdapter{
		conn: c,
	}
}

func (pa *PostgresAdapter) Create() {}

func (pa *PostgresAdapter) Read() {}

func (pa *PostgresAdapter) Ping(ctx context.Context) (bool, error) {
	err := pa.conn.PingContext(ctx)
	if err != nil {
		return false, err
	}

	return true, nil
}
