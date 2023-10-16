package transaction

import (
	"fmt"
	"time"
)

type Operation int

const (
	OpCompraAVista Operation = iota + 1
	OpCompraParcelada
	OpSaque
	OpPagamento
)

type Transaction struct {
	Id          int       `json:"id"`
	AccountId   int       `json:"account_id"`
	OperationId Operation `json:"operation_id"`
	Amount      float32   `json:"amount"`
	CreatedAt   time.Time `json:"-"`
}

func (trx *Transaction) Valid() bool {
	// A Transaction should not have it's amount equal to zero.
	if trx.Amount == 0 {
		return false
	}
	// A Transaction should always be linked to an Account. There are no Accounts with the Id == zero.
	if trx.AccountId == 0 {
		return false
	}
	switch trx.OperationId {
	case OpCompraAVista, OpCompraParcelada, OpSaque:
		// A Transaction with OperationId == OpCompraAVista, OpCompraParcelada or OpSaque
		// should always have a negative value.
		if trx.Amount > 0 {
			return false
		}
	case OpPagamento:
		// A Transaction with OperationId == OpPagamento should always have a positive value.
		if trx.Amount < 0 {
			return false
		}
	}

	return true
}

func (trx *Transaction) String() string {
	return fmt.Sprintf("Transaction [ Id: %d, Account Id: %d, Operation Id: %d, Amount: %f.2, Created At: %v ]", trx.Id, trx.AccountId, trx.OperationId, trx.Amount, trx.CreatedAt)
}
