package model

type date struct {
	value string
}

//Date interface representing the date struct
type Date interface {
	GetValue() string
}

//NewDate create a new date struct
func NewDate(value string) *date {
	return &date{
		value: value,
	}
}

//GetValue returns the value of date
func (date *date) GetValue() string {
	return date.value
}
