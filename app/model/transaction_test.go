package model

import (
	"testing"
)

func TestNewTransactionValid(t *testing.T) {
	values := []int{1, 2, 3, 4}

	for key, value := range values {
		var id, account, operationType, amount interface{}
		var amountByOperatonType float64
		documentModel, _ := NewDocument("99407901041")
		accountModel := NewAccount(documentModel)
		operationTypeModel, _ := NewOperationType(value)
		amountModel, _ := NewAmount(10.0)

		transactionModel := NewTransaction(accountModel, operationTypeModel, amountModel)

		id = transactionModel.GetId()
		account = transactionModel.GetAccount()
		operationType = transactionModel.GetOperationType()
		amount = transactionModel.GetAmount()

		amountByOperatonType = amountModel.GetValue()

		if value != 4 {
			amountByOperatonType = amountModel.GetValue() * -1
		}

		if amountByOperatonType != transactionModel.GetAmountValueByOperationType() {
			t.Errorf("\nTest at position [%d].\nInvalid amount by operation type: '%f'", key, amountByOperatonType)
		}

		if _, ok := id.(ID); !ok {
			t.Errorf("\nTest at position [%d].\nInvalid ID type: '%T'", key, id)
		}

		if _, ok := account.(Account); !ok {
			t.Errorf("\nTest at position [%d].\nInvalid account type: '%T'", key, id)
		}

		if _, ok := operationType.(OperationType); !ok {
			t.Errorf("\nTest at position [%d].\nInvalid operationType type: '%T'", key, id)
		}

		if _, ok := amount.(Amount); !ok {
			t.Errorf("\nTest at position [%d].\nInvalid operationType type: '%T'", key, id)
		}
	}
}
