package database

import (
	"context"
)

type Utils struct {
	db Adapter
}

func NewUtils(adapter Adapter) *Utils {
	return &Utils{
		db: adapter,
	}
}

func (u *Utils) PingDatabase(ctx context.Context) (bool, error) {
	return u.db.Ping(ctx)
}
