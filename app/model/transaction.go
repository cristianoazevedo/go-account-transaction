package model

type transaction struct {
	id            ID
	account       Account
	operationType OperationType
	amount        Amount
	eventDate     Date
}

//Transaction interface representing the transaction struct
type Transaction interface {
	GetID() ID
	GetAccount() Account
	GetOperationType() OperationType
	GetAmount() Amount
	GetAmountValueByOperationType() float64
}

//NewTransaction Create a new transaction struct
func NewTransaction(account Account, operationType OperationType, amount Amount) (Transaction, error) {
	err := checkAccountCreditLimt(account, operationType, amount)

	if err != nil {
		return nil, err
	}

	transactionModel := &transaction{
		id:            NewID(),
		account:       account,
		operationType: operationType,
		amount:        amount,
		eventDate:     NewDate(),
	}

	account.NewCreditLimit(transactionModel.GetAmountValueByOperationType())

	return transactionModel, nil
}

func checkAccountCreditLimt(account Account, operationType OperationType, amount Amount) error {
	if operationType.GetValue() != Payment && !account.HasCreditLimit(amount.GetValue()) {
		return NewDomainError("limit insufient")
	}

	return nil
}

//GetId return the struct ID
func (transaction *transaction) GetID() ID {
	return transaction.id
}

//GetAccount return the struct Account
func (transaction *transaction) GetAccount() Account {
	return transaction.account
}

//GetOperationType return the struct Operation Type
func (transaction *transaction) GetOperationType() OperationType {
	return transaction.operationType
}

//GetAmount return the struct Amount
func (transaction *transaction) GetAmount() Amount {
	return transaction.amount
}

//GetAmountValueByOperationType return the amount by operation type
func (transaction *transaction) GetAmountValueByOperationType() float64 {
	if transaction.operationType.GetValue() == Payment {
		return transaction.amount.GetValue()
	}

	return transaction.amount.GetNegativeValue()
}
