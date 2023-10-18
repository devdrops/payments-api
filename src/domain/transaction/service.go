package transaction

import (
	"context"
)

type DischargeService struct {
	TrxRep *TransactionRepository
}

func NewDischargeService() *DischargeService {
	return &DischargeService{}
}

func (ds *DischargeService) ExecuteDischarges(ctx context.Context) {
	// Pick Operations based on type
	operations := []Operation{OpCompraAVista, OpCompraParcelada, OpSaque}
	transactions, err := ds.TrxRep.GetBalanceByOperations(ctx, operations)
	if err != nil {
		return err
	}


}
