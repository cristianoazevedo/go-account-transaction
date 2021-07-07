package model

type operationType struct {
	value int
}

const (
	//CashPurchase has the value 1
	CashPurchase = iota + 1
	//PurchaseParceled has the value 2
	PurchaseParceled
	//Withdraw has the value 3
	Withdraw
	//Payment has the value 4
	Payment
)

//OperationType interface representing the operationType struct
type OperationType interface {
	GetValue() int
}

//NewOperationType create a new operationType struct
func NewOperationType(value int) (OperationType, error) {
	if value < CashPurchase || value > Payment {
		return nil, NewDomainError("operation type invalid")
	}

	return &operationType{value: value}, nil
}

//GetValue returns the value of operation type
func (operationType *operationType) GetValue() int {
	return operationType.value
}
