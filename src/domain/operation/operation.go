package operation

import (
	"fmt"
	"time"
)

type Description string

const (
	OpCompraAVista    Description = "COMPRA_A_VISTA"
	OpCompraParcelada Description = "COMPRA_PARCELADA"
	OpSaque           Description = "SAQUE"
	OpPagamento       Description = "PAGAMENTO"
)

type Operation struct {
	Id          int
	Description Description
	CreatedAt   time.Time
}

func (op *Operation) String() string {
	return fmt.Sprintf("Operation [ Id: %d, Description: %s, Created At: %v ]", op.Id, op.Description, op.CreatedAt)
}
