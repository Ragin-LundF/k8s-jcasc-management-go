package validator

import (
	"testing"
)

func TestValidateIP(t *testing.T) {
	var ip = "1.2.3.4"
	err := ValidateIP(ip)

	if err != nil {
		t.Error("Failed. Validate has thrown an error.")
	} else {
		t.Log("Success. IP was detected correct.")
	}
}

func TestValidateIPWrongNumber(t *testing.T) {
	var ip = "1.2.321.4"
	err := ValidateIP(ip)

	if err != nil {
		t.Log("Success. Wrong IP was detected correct.")
	} else {
		t.Error("Failed. Validate has thrown no error.")
	}
}

func TestValidateIPWithThreeNumbers(t *testing.T) {
	var ip = "1.2.321"
	err := ValidateIP(ip)

	if err != nil {
		t.Log("Success. Wrong IP with 3 numbers was detected correct.")
	} else {
		t.Error("Failed. Validate has thrown no error.")
	}
}

func TestValidateIPWithSimpleString(t *testing.T) {
	var ip = "test"
	err := ValidateIP(ip)

	if err != nil {
		t.Log("Success. String instead of IP was detected correct.")
	} else {
		t.Error("Failed. Validate has thrown no error.")
	}
}
