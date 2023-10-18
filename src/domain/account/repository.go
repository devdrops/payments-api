package account

import (
	"context"
	"fmt"

	"payments-api/internal/database"
)

const table = "accounts"

type AccountRepository struct {
	db database.Adapter
}

func NewRepository(adapter database.Adapter) *AccountRepository {
	return &AccountRepository{
		db: adapter,
	}
}

func (rep *AccountRepository) Create(ctx context.Context, doc string) error {
	columns := []string{"document"}

	err := rep.db.Insert(ctx, table, columns, doc)
	if err != nil {
		return err
	}

	return nil
}

func (rep *AccountRepository) Get(ctx context.Context, id int) (Account, error) {
	acc := Account{}
	columns := []string{"id", "document"}
	cond := fmt.Sprintf("WHERE id = %d", id)
	conditions := []string{cond}

	row := rep.db.GetOne(ctx, table, columns, conditions)
	err := row.Scan(&acc.Id, &acc.Document)
	if err != nil {
		return acc, err
	}

	return acc, nil
}
