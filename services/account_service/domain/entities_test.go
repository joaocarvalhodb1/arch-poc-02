package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewProfileRule(t *testing.T) {
	var testCases = []struct {
		testName         string
		inputName        string
		inputEmail       string
		inputCellPhone   string
		inputAccoutType  string
		expectedResponse string
	}{
		{testName: "ok", inputName: "John doe", inputEmail: "john@email.com", inputCellPhone: "+55 999999999", inputAccoutType: "LEAD", expectedResponse: ""},
		{testName: "name empty", inputName: "", inputEmail: "john@email.com", inputCellPhone: "+55 999999999", inputAccoutType: "LEAD", expectedResponse: "name is required"},
		{testName: "email empty", inputName: "John doe", inputEmail: "", inputCellPhone: "+55 999999999", inputAccoutType: "LEAD", expectedResponse: "email is required"},
		{testName: "cell phone empty", inputName: "John doe", inputEmail: "john@email.com", inputCellPhone: "", inputAccoutType: "LEAD", expectedResponse: "cell phone is required"},
		{testName: "account type empty", inputName: "John doe", inputEmail: "john@email.com", inputCellPhone: "+55 999999999", inputAccoutType: "", expectedResponse: "account type is required"},
	}
	for _, test := range testCases {
		account, err := NewAccount(test.inputName, test.inputEmail, test.inputCellPhone, 0)
		if err == nil {
			assert.NotNil(t, account)
			assert.Equal(t, "John doe", account.Name)
			assert.Equal(t, "john@email.com", account.Email)
			assert.Equal(t, "+55 999999999", account.CellPhone)
		} else {
			assert.NotNil(t, err)
			assert.Error(t, err, test.expectedResponse)
		}
	}
}
