package database

import (
	"context"

	// TODO: solve this abstraction failure
	"database/sql"
)

type Adapter interface {
	Insert(ctx context.Context, table string, columns []string, values ...any) error
	GetOne(ctx context.Context, table string, columns []string, filters []string) *sql.Row
	GetMany(ctx context.Context, table string, columns []string, filters []string) (*sql.Rows, error)
	Ping(ctx context.Context) (bool, error)
}
