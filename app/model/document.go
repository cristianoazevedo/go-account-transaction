package model

import "github.com/Nhanderu/brdoc"

type document struct {
	value string
}

//Document interface representing the document struct
type Document interface {
	GetValue() string
}

//NewDocument create a new document struct
func NewDocument(value string) (Document, error) {
	if !brdoc.IsCPF(value) {
		return nil, NewDomainError("document invalid")
	}

	return &document{value: value}, nil
}

//GetValue returns the value of document
func (document *document) GetValue() string {
	return document.value
}
