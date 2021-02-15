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

func TestValidateDomain(t *testing.T) {
	var ip = "www.domain.tld"
	err := ValidateIP(ip)

	assert.NoError(t, err)
}

func TestValidateDomainWithNumbers(t *testing.T) {
	var ip = "test--int-123.test.k8s.cluster.project1.domain.int"
	err := ValidateIP(ip)

	assert.NoError(t, err)
}

func TestValidateDomainWithWrongDomain(t *testing.T) {
	var ip = "domain"
	err := ValidateIP(ip)

	assert.Error(t, err)
}

func TestValidateIPWithSimpleString(t *testing.T) {
	var ip = "test"
	err := ValidateIP(ip)

	assert.Error(t, err)
}
