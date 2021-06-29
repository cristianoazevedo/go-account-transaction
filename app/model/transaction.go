package model

type transaction struct {
	id            *id
	account       account
	operationType OperationType
	amount        amount
	eventDate     date
}

func NewTransaction(account account) *transaction {
	return &transaction{
		id:            NewID(),
		account:       account,
		operationType: 0,
		amount: amount{
			value: 0,
		},
		eventDate: date{},
	}
}
