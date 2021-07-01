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

type OperationType interface {
	GetValue() int
}

func NewOperationType(value int) (*operationType, error) {
	if value < CashPurchase || value > Payment {
		return nil, NewDomainError("Operation type invaild")
	}

	return &operationType{value: value}, nil
}

func (operationType *operationType) GetValue() int {
	return operationType.value
}
