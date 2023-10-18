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
	columns := []string{"account_id", "operation_id", "amount", "balance"}

	err := rep.db.Insert(ctx, table, columns, aId, oId, a, a)
	if err != nil {
		return err
	}

	return nil
}

func (rep *TransactionRepository) GetBalanceByOperations(ctx context.Context, operations []Operation) ([]Transaction, error) {
	columns := []string{"amount", "balance"}
	values := make([]string, len(values))
	for i := 0; i < len(values); i++ {
		v[i] = fmt.Sprintf("$%d", i+1)
	}
	filters := []string{"WHERE operation_id"}

	rows, err := rep.db.GetMany(ctx, table, columns, )
}
