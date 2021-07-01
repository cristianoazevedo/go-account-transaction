package model

import (
	"testing"
)

func TestOperationTypesValid(t *testing.T) {
	operationTypes := []int{1, 2, 3, 4}

	for key, value := range operationTypes {
		_, err := NewOperationType(value)

		if err != nil {
			t.Errorf(`\nTest at position [%d].\nError: '%v'`, key, err)
		}
	}
}
