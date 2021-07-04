package model

import "math"

type amount struct {
	value float64
}

//Amount interface representing the amount struct
type Amount interface {
	GetValue() float64
	GetNegativeValue() float64
}

//NewAmount create a new amount struct
func NewAmount(value float64) (Amount, error) {
	if value <= 0 || value > math.MaxFloat64 {
		return nil, NewDomainError("amount not allowed")
	}

	return &amount{value: value}, nil
}

//GetValue returns the value of amount
func (amount *amount) GetValue() float64 {
	return amount.value
}

//GetNegativeValue returns the negative value of amount
func (amount *amount) GetNegativeValue() float64 {
	return amount.value * -1
}
