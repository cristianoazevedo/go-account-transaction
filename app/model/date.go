package model

type date struct {
	value string
}

type Date interface {
	GetValue() string
}

func NewDate(value string) *date {
	return &date{
		value: value,
	}
}

func (date *date) GetValue() string {
	return date.value
}
