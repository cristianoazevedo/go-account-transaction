package model

import "github.com/Nhanderu/brdoc"

type document struct {
	value string
}

func NewDocument(value string) (*document, error) {
	if !brdoc.IsCPF(value) {
		return nil, NewDomainError("Document invalid")
	}

	return &document{value: value}, nil
}

type Document interface {
	GetValue() string
}

func (document *document) GetValue() string {
	return document.value
}
