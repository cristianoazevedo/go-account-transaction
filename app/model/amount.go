package model

import "math"

type amount struct {
	value float64
}

type Amount interface {
	GetValue() float64
}

func NewAmount(value float64) (*amount, error) {
	if value <= 0 || value > math.MaxFloat64 {
		return nil, NewDomainError("amount not allowed")
	}

	return &amount{value: value}, nil
}

func (amount *amount) GetValue() float64 {
	return amount.value
}
