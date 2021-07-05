package model

const limit float64 = 500

type availableCreditLimit struct {
	value float64
}

type AvailableCreditLimit interface {
	GetValue() float64
}

func NewAvailableCreditLimit() AvailableCreditLimit {
	return &availableCreditLimit{value: limit}
}

func BuildAvailableCreditLimit(value float64) (AvailableCreditLimit, error) {
	if value <= 0 {
		return nil, NewDomainError("limit invalid")
	}

	return &availableCreditLimit{value: value}, nil
}

func (creditLimit *availableCreditLimit) GetValue() float64 {
	return creditLimit.value
}
