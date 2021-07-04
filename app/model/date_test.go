package model

import "testing"

func TestNewDateValid(t *testing.T) {
	var value interface{}
	dateModel := NewDate()
	value = dateModel.GetValue()

	if _, ok := value.(string); !ok {
		t.Errorf("\nInvalid type: '%T'", value)
	}
}

func TestBuildDateValid(t *testing.T) {
	dates := []string{
		"2021-07-02 16:39:08",
		"2021-08-12 00:00:00",
		"2020-11-15 09:45:08",
		"2019-03-09 12:00:08",
	}

	for key, date := range dates {
		var value interface{}
		dataModel := BuildDate(date)
		value = dataModel.GetValue()

		if _, ok := value.(string); !ok {
			t.Errorf("\nTest at position [%d].\nInvalid type: '%T'", key, value)
		}

		if dataModel.GetValue() != date {
			t.Errorf("\nTest at position [%d].\nInvalid value: '%s'", key, value)
		}
	}
}
