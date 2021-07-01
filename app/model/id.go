package model

import "github.com/google/uuid"

type id struct {
	value string
}

type ID interface {
	GetValue() string
}

func NewID() *id {
	return &id{value: uuid.New().String()}
}

func BuildId(value string) (*id, error) {
	valueParsed, err := uuid.Parse(value)

	if err != nil {
		return nil, err
	}

	return &id{value: valueParsed.String()}, nil
}

func (id *id) GetValue() string {
	return id.value
}
