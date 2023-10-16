package transaction_test

import (
	"testing"

	"payments-api/src/domain/transaction"
)

func TestValid(t *testing.T) {
	data := []struct {
		amount   float32
		accId    int
		opId     transaction.Operation
		expected bool
	}{
		{amount: 14.06, accId: 1, opId: transaction.OpPagamento, expected: true},
		{amount: 101.10, accId: 2, opId: transaction.OpCompraAVista, expected: false},
		{amount: 0, accId: 3, opId: transaction.OpPagamento, expected: false},
		{amount: -6.66, accId: 4, opId: transaction.OpCompraParcelada, expected: true},
		{amount: -100.01, accId: 5, opId: transaction.OpCompraAVista, expected: true},
		{amount: -100.01, accId: 0, opId: transaction.OpCompraAVista, expected: false},
		{amount: 123.45, accId: 7, opId: transaction.OpPagamento, expected: true},
		{amount: -100.06, accId: 1, opId: transaction.OpPagamento, expected: false},
	}

	for _, elem := range data {
		trx := transaction.Transaction{
			Amount:      elem.amount,
			AccountId:   elem.accId,
			OperationId: elem.opId,
		}

		if trx.Valid() != elem.expected {
			t.Errorf("Test failed: %v\n", elem)
		}
	}
}
