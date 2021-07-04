package model

import "time"

type date struct {
	value string
}

const layout = "2006-01-02 15:04:05"

//Date interface representing the date struct
type Date interface {
	GetValue() string
}

//NewDate create a new date struct
func NewDate() Date {
	currentTimestamp := time.Now()

	return &date{value: currentTimestamp.Format(layout)}
}

//BuildDate create a new date struct, with parameters passed
func BuildDate(value string) Date {
	valueParsed, _ := time.Parse(layout, value)

	return &date{value: valueParsed.Format(layout)}
}

//GetValue returns the value of date
func (date *date) GetValue() string {
	return date.value
}
