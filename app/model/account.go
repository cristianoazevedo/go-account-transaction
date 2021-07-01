package model

type account struct {
	id        ID
	createdAt Date
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
		createdAt: NewDate("today"),
		document:  document,
	}
}

func BuildAccount(id string, document string, createdAt string) (*account, error) {
	idBuilded, err := BuildId(id)

	if err != nil {
		return nil, err
	}

	newDocument, err := NewDocument(document)

	if err != nil {
		return nil, err
	}

	return &account{id: idBuilded, document: newDocument, createdAt: NewDate(createdAt)}, nil
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
