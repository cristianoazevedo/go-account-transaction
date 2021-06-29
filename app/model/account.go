package model

type account struct {
	id        *id
	createdAt *date
	document  Document
}

type Account interface {
	GetId() ID
	GetDocument() Document
	GetCreatedAt() Date
}

func NewAccount(document Document) *account {
	return &account{
		id:        NewID(),
		createdAt: NewDate(),
		document:  document,
	}
}

func (account *account) GetId() ID {
	return account.id
}

func (account *account) GetDocument() Document {
	return account.document
}

func (account *account) GetCreatedAt() Date {
	return account.createdAt
}
