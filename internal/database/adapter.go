package database

import (
	"context"
)

type Adapter interface {
	Create(ctx context.Context, t string, args ...any) error
	Read(ctx context.Context, t string, args ...any)
	Ping(ctx context.Context) (bool, error)
}
