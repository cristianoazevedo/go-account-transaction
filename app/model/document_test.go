package model

import "testing"

func TestDocumentNumberValid(t *testing.T) {
	values := []string{
		"99407901041",
		"60398738092",
		"06430068005",
		"61758351071",
	}

	for key, value := range values {
		documentNumberModel, err := NewDocument(value)

		if err != nil {
			t.Errorf("\nTest at position [%d].\nAn error '%s' was not expected", key, err)
			continue
		}

		if documentNumberModel.GetValue() != value {
			t.Errorf("\nTest at position [%d].\nInvalid Value: '%v'", key, value)
		}
	}
}

func TestDocumentNumberInvalid(t *testing.T) {
	values := []string{
		"00000000001",
		"00000000012",
		"00000000123",
		"00000001234",
	}

	for key, value := range values {
		_, err := NewDocument(value)

		if err == nil {
			t.Errorf("\nTest at position [%d].\nExpected error", key)
		}
	}
}
