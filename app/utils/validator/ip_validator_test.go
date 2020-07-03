package validator

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValidateIP(t *testing.T) {
	var ip = "5.2.3.4"
	err := ValidateIP(ip)

	assert.NoError(t, err)
}

func TestValidateIPAlreadyExisting(t *testing.T) {
	var ip = "1.2.3.4"
	err := ValidateIP(ip)

	assert.Error(t, err)
}

func TestValidateIPWrongNumber(t *testing.T) {
	var ip = "5.2.321.4"
	err := ValidateIP(ip)

	assert.Error(t, err)
}

func TestValidateIPWithThreeNumbers(t *testing.T) {
	var ip = "5.2.321"
	err := ValidateIP(ip)

	assert.Error(t, err)
}

func TestValidateIPWithSimpleString(t *testing.T) {
	var ip = "test"
	err := ValidateIP(ip)

	assert.Error(t, err)
}
