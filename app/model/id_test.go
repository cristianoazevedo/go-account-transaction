package model

import (
	"testing"
)

func TestNewIdValid(t *testing.T) {
	var value interface{}
	idModel := NewID()
	value = idModel.GetValue()

	if _, ok := value.(string); !ok {
		t.Errorf("\nInvalid type: '%T'", value)
	}
}

func TestBuildIdValid(t *testing.T) {
	values := []string{
		"99c49b65-cc11-487f-864d-55dbb6c90a67",
		"83ef78ca-478b-49fd-9514-d285f146969b",
		"dd8cc5e9-d6cd-4ba1-ae08-6b0137e51cc4",
		"9e3b6419-e696-4a1b-8c14-8a66dcaf6d15",
		"69e94b53-da7f-4156-bb32-c923dc3fd682",
		"7019db39-c9e1-4d95-a9dc-a9d8741a6dbf",
		"0d39290a-a074-4765-9fc7-d37ade6e3b86",
	}

	for key, value := range values {
		idModel, err := BuildID(value)

		if err != nil {
			t.Errorf("\nTest at position [%d].\nAn error '%s' was not expected", key, err)
			break
		}

		if idModel.GetValue() != value {
			t.Errorf("\nTest at position [%d].\nInvalid Value: '%v'", key, value)
		}
	}
}

func TestBuildIdInalid(t *testing.T) {
	values := []string{
		"99c49b65-cc11-487f-864d-55dbb6c90a6",
		"83ef78ca-478b-49fd-9514-d285f146969b1",
		"dd8cc5e9-d6cd-4ba1-ae08-6b0137e51",
		"9e3b6419-e696-4a1b-8c14",
		"69e94b53-da7f-4156",
		"7019db39-c9e1",
		"0d39290a-a074",
	}

	for key, value := range values {
		_, err := BuildID(value)

		if err == nil {
			t.Errorf("\nTest at position [%d].\nExpected error", key)
		}
	}
}
