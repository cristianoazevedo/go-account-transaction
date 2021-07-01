package model

import (
	"testing"
)

func TestOperationTypesValid(t *testing.T) {
	values := []int{1, 2, 3, 4}

	for key, value := range values {
		operationTypeModel, err := NewOperationType(value)

		if err != nil {
			t.Errorf("\nTest at position [%d].\nAn error '%s' was not expected", key, err)
			continue
		}

		if operationTypeModel.GetValue() != value {
			t.Errorf("\nTest at position [%d].\nInvalid Value: '%v'", key, value)
		}
	}
}

func TestOperationTypesInvalid(t *testing.T) {
	values := []int{0, 5, 6, 7}

	for key, value := range values {
		_, err := NewOperationType(value)

		if err == nil {
			t.Errorf("\nTest at position [%d].\nExpected error", key)
		}
	}
}
