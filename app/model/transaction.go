package model

type transaction struct {
	id            ID
	account       Account
	operationType OperationType
	amount        Amount
	eventDate     Date
}

type Transaction interface {
	GetId() ID
	GetAccount() Account
	GetOperationType() OperationType
	GetAmount() Amount
	GetAmountValueByOperationType() float64
}

func NewTransaction(account Account, operationType OperationType, amount Amount) *transaction {
	return &transaction{
		id:            NewID(),
		account:       account,
		operationType: operationType,
		amount:        amount,
		eventDate:     NewDate("now"),
	}
}

func (transaction *transaction) GetId() ID {
	return transaction.id
}

func (transaction *transaction) GetAccount() Account {
	return transaction.account
}

func (transaction *transaction) GetOperationType() OperationType {
	return transaction.operationType
}

func (transaction *transaction) GetAmount() Amount {
	return transaction.amount
}

func (transaction *transaction) GetAmountValueByOperationType() float64 {
	if transaction.operationType.GetValue() == Payment {
		return transaction.amount.GetValue()
	}

	return transaction.amount.GetValue() * -1
}
