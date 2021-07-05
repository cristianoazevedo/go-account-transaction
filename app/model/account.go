package model

type account struct {
	id                   ID
	createdAt            Date
	document             Document
	availableCreditLimit AvailableCreditLimit
}

//Account interface representing the account struct
type Account interface {
	GetID() ID
	GetDocument() Document
	GetCreatedAt() Date
	GetAvailableCreditLimit() AvailableCreditLimit
	NewCreditLimit(newLimit float64)
	HasCreditLimit(limit float64) bool
}

//NewAccount create a new account struct
func NewAccount(document Document, creditLimit AvailableCreditLimit) Account {
	return &account{
		id:                   NewID(),
		createdAt:            NewDate(),
		document:             document,
		availableCreditLimit: creditLimit,
	}
}

//BuildAccount create a new account struct, with parameters passed
func BuildAccount(id string, document string, createdAt string, creditLimit float64) (Account, error) {
	idBuilded, err := BuildID(id)

	if err != nil {
		return nil, err
	}

	newDocument, err := NewDocument(document)

	if err != nil {
		return nil, err
	}

	newCreditLimit, err := BuildAvailableCreditLimit(creditLimit)

	if err != nil {
		return nil, err
	}

	return &account{id: idBuilded, document: newDocument, createdAt: BuildDate(createdAt), availableCreditLimit: newCreditLimit}, nil
}

//GetId return the struct ID
func (account *account) GetID() ID {
	return account.id
}

//GetDocument return the struct Document
func (account *account) GetDocument() Document {
	return account.document
}

//GetCreatedAt return the struct Date
func (account *account) GetCreatedAt() Date {
	return account.createdAt
}

func (account *account) GetAvailableCreditLimit() AvailableCreditLimit {
	return account.availableCreditLimit
}

func (account *account) NewCreditLimit(value float64) {
	newLimt := account.availableCreditLimit.GetValue() + value
	creditLimit, _ := BuildAvailableCreditLimit(newLimt)
	account.availableCreditLimit = creditLimit
}

func (account *account) HasCreditLimit(limit float64) bool {
	return account.availableCreditLimit.GetValue() >= limit
}
