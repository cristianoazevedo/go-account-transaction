package model

import "github.com/google/uuid"

type id struct {
	value string
}

//ID interface representing the id struct
type ID interface {
	GetValue() string
}

//NewID create a new id struct
func NewID() *id {
	return &id{value: uuid.New().String()}
}

//BuildID create a new id struct, with parameters passed
func BuildID(value string) (*id, error) {
	valueParsed, err := uuid.Parse(value)

	if err != nil {
		return nil, err
	}

	return &id{value: valueParsed.String()}, nil
}

//GetValue returns the value of id
func (id *id) GetValue() string {
	return id.value
}
