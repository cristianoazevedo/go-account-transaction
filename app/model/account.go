package model

type account struct {
	id        ID
	createdAt Date
	document  Document
}

//Account interface representing the account struct
type Account interface {
	GetID() ID
	GetDocument() Document
	GetCreatedAt() Date
}

//NewAccount create a new account struct
func NewAccount(document Document) *account {
	return &account{
		id:        NewID(),
		createdAt: NewDate("today"),
		document:  document,
	}
}

//BuildAccount create a new account struct, with parameters passed
func BuildAccount(id string, document string, createdAt string) (*account, error) {
	idBuilded, err := BuildID(id)

	if err != nil {
		return nil, err
	}

	newDocument, err := NewDocument(document)

	if err != nil {
		return nil, err
	}

	return &account{id: idBuilded, document: newDocument, createdAt: NewDate(createdAt)}, nil
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
