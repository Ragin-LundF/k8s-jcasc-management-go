package config

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseIpConfigurationLine(t *testing.T) {
	var namespaceToCheck = "mynamespace"
	var ipToCheck = "1.2.3.4"
	var line = namespaceToCheck + " " + ipToCheck
	namespace, ip := parseIPConfigurationLine(line)

	assert.Equal(t, namespaceToCheck, namespace)
	assert.Equal(t, ipToCheck, ip)
}

func TestParseIpConfigurationWithEqual(t *testing.T) {
	var namespaceToCheck = "mynamespace"
	var ipToCheck = "1.2.3.4"
	var line = namespaceToCheck + "=" + ipToCheck
	namespace, ip := parseIPConfigurationLine(line)

	assert.Equal(t, namespaceToCheck, namespace)
	assert.Equal(t, ipToCheck, ip)
}

func TestParseIpConfigurationWithSpaces(t *testing.T) {
	var namespaceToCheck = "mynamespace"
	var ipToCheck = "1.2.3.4"
	var line = namespaceToCheck + " = " + ipToCheck
	namespace, ip := parseIPConfigurationLine(line)

	assert.Equal(t, namespaceToCheck, namespace)
	assert.Equal(t, ipToCheck, ip)
}

func TestParseIpConfigurationWithInvalidLine(t *testing.T) {
	var namespaceToCheck = "mynamespace"
	var line = namespaceToCheck
	namespace, ip := parseIPConfigurationLine(line)

	assert.Equal(t, "", namespace)
	assert.Equal(t, "", ip)
}
