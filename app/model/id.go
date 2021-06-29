package model

import "github.com/google/uuid"

type id struct {
	value string
}

type ID interface {
	GetValue() string
}

func NewID() *id {
	value, _ := uuid.NewRandom()

	return &id{value: value.String()}
}

func (id *id) GetValue() string {
	return id.value
}
