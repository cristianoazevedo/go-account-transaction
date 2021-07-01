package model

import (
	"testing"
)

func TestNewAmountValid(t *testing.T) {
	values := []float64{
		.1,
		1,
		1.23,
		10.99,
		99.99,
		100,
		123.33,
		999.99,
		1512.66,
		20500.78,
	}

	for key, value := range values {
		amountModel, err := NewAmount(value)

		if err != nil {
			t.Errorf("\nTest at position [%d].\nAn error '%s' was not expected", key, err)
			continue
		}

		if amountModel.GetValue() != value {
			t.Errorf("\nTest at position [%d].\nInvalid Value: '%v'", key, value)
		}
	}
}

func TestNewAmountInvalid(t *testing.T) {
	_, err := NewAmount(0)

	if err == nil {
		t.Errorf("\nTest expected error")
	}
}
