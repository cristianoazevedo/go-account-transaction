package model

import "testing"

func TestNewAcccountValid(t *testing.T) {
	values := []string{
		"99407901041",
		"60398738092",
		"06430068005",
		"61758351071",
	}

	for key, value := range values {
		var id, document, createdAt interface{}
		documentNumberModel, _ := NewDocument(value)
		accountModel := NewAccount(documentNumberModel)
		id = accountModel.GetID()
		document = accountModel.GetDocument()
		createdAt = accountModel.GetCreatedAt()

		if _, ok := id.(ID); !ok {
			t.Errorf("\nTest at position [%d].\nInvalid type: '%T'", key, value)
		}

		if _, ok := document.(Document); !ok {
			t.Errorf("\nTest at position [%d].\nInvalid type: '%T'", key, value)
		}

		if _, ok := createdAt.(Date); !ok {
			t.Errorf("\nTest at position [%d].\nInvalid type: '%T'", key, value)
		}
	}
}

func TestBuildAccountValid(t *testing.T) {
	values := []struct {
		id             string
		documentNumber string
		createdAt      string
	}{
		{
			"99c49b65-cc11-487f-864d-55dbb6c90a67",
			"99407901041",
			"2019-03-09 12:00:08",
		},
		{
			"83ef78ca-478b-49fd-9514-d285f146969b",
			"60398738092",
			"2019-03-09 12:00:08",
		},
	}

	for key, value := range values {
		accountModel, err := BuildAccount(value.id, value.documentNumber, value.createdAt)

		if err != nil {
			t.Errorf("\nTest at position [%d].\nAn error '%s' was not expected", key, err)
			break
		}

		if accountModel.GetID().GetValue() != value.id {
			t.Errorf("\nTest at position [%d].\nInvalid ID value: '%v'", key, value.id)
		}

		if accountModel.GetDocument().GetValue() != value.documentNumber {
			t.Errorf("\nTest at position [%d].\nInvalid document value: '%v'", key, value.documentNumber)
		}

		if accountModel.GetCreatedAt().GetValue() != value.createdAt {
			t.Errorf("\nTest at position [%d].\nInvalid date value: '%v'", key, value.createdAt)
		}
	}
}

func TestBuildAccountInvalid(t *testing.T) {
	values := []struct {
		id             string
		documentNumber string
		createdAt      string
	}{
		{
			"99c49b65-cc11-487f-864d",
			"99407901041",
			"2019-03-09 12:00:08",
		},
		{
			"99c49b65-cc11-487f-864d-55dbb6c90a67",
			"00000000002",
			"2019-03-09 12:00:08",
		},
	}

	for key, value := range values {
		_, err := BuildAccount(value.id, value.documentNumber, value.createdAt)

		if err == nil {
			t.Errorf("\nTest at position [%d].\nExpected error", key)
		}
	}
}
