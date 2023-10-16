package account_test

import (
	"testing"

	"payments-api/src/domain/account"
)

// Test if the validation for an Account is working as expected
func TestValid(t *testing.T) {
	data := []struct {
		doc      string
		expected bool
	}{
		{doc: "12345678900", expected: true},
		{doc: "1234567890000", expected: true},
		{doc: "ABC12364578900", expected: true},
		{doc: "123.456.789-00", expected: false},
		{doc: "12.345.678/0001-99", expected: false},
		{doc: "12 45 78900 ", expected: false},
		{doc: "", expected: false},
		{doc: "            ", expected: false},
		{doc: "112223334", expected: false},
		{doc: "aaB12333444", expected: true},
		{doc: "111AAA222BBB", expected: true},
	}

	for _, elem := range data {
		acc := account.Account{
			Document: elem.doc,
		}

		if acc.Valid() != elem.expected {
			t.Errorf("Test failed: %v\n", elem)
		}
	}
}
