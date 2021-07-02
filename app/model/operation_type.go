package model

type operationType struct {
	value int
}

const (
	CashPurchase = iota + 1
	PurchaseParceled
	Withdraw
	Payment
)

//OperationType interface representing the operationType struct
type OperationType interface {
	GetValue() int
}

//NewOperationType create a new operationType struct
func NewOperationType(value int) (*operationType, error) {
	if value < CashPurchase || value > Payment {
		return nil, NewDomainError("operation type invaild")
	}

	return &operationType{value: value}, nil
}

//GetValue returns the value of operation type
func (operationType *operationType) GetValue() int {
	return operationType.value
}
