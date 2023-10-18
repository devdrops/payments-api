package database

import (
	"context"
)

type Adapter interface {
	Insert(ctx context.Context, table string, columns []string, values ...any) error
	GetOne(ctx context.Context, table string, columns []string, filters []string, ent Entity) (Entity, error)
	Ping(ctx context.Context) (bool, error)
}
