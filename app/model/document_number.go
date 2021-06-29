package model

type document struct {
	value string
}

func NewDocument(value string) *document {
	return &document{
		value: value,
	}
}

type Document interface {
	GetValue() string
}

func (document *document) GetValue() string {
	return document.value
}
