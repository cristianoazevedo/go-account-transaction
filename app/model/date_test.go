package model

import "testing"

func TestNewDateValid(t *testing.T) {
	var value interface{}
	dataModel := NewDate("today")
	value = dataModel.GetValue()

	if _, ok := value.(string); !ok {
		t.Errorf("\nInvalid type: '%T'", value)
	}
}
