package model

type date struct {
	value string
}

type Date interface {
	GetValue() string
}

func NewDate() *date {
	return &date{
		value: "",
	}
}

func (date *date) GetValue() string {
	return date.value
}
