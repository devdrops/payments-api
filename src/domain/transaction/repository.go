package transaction

import (
	"context"

	"payments-api/internal/database"
)

const table = "transactions"

type TransactionRepository struct {
	db database.Adapter
}

func NewRepository(adapter database.Adapter) *TransactionRepository {
	return &TransactionRepository{
		db: adapter,
	}
}

func (rep *TransactionRepository) Create(ctx context.Context, aId int, oId Operation, a float32) error {
	columns := []string{"account_id", "operation_id", "amount"}

	err := rep.db.Insert(ctx, table, columns, aId, oId, a)
	if err != nil {
		return err
	}

	return nil
}
