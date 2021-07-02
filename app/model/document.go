package model

import "github.com/Nhanderu/brdoc"

type document struct {
	value string
}

//NewDocument create a new document struct
func NewDocument(value string) (*document, error) {
	if !brdoc.IsCPF(value) {
		return nil, NewDomainError("Document invalid")
	}

	return &document{value: value}, nil
}

//Document interface representing the document struct
type Document interface {
	GetValue() string
}

//GetValue returns the value of document
func (document *document) GetValue() string {
	return document.value
}
