package config

import (
	"github.com/stretchr/testify/assert"
	"k8s-management-go/app/models"
	"testing"
)

// Read configuration from k8s-management config file
func TestParseConfigurationLine(t *testing.T) {
	key, value := parseConfigurationLine("KEY=MYVALUE")

	assert.Equal(t, "KEY", key)
	assert.Equal(t, "MYVALUE", value)
}

func TestParseConfigurationLineWithDoubleQuotes(t *testing.T) {
	key, value := parseConfigurationLine("KEY=\"MYVALUE\"")

	assert.Equal(t, "KEY", key)
	assert.Equal(t, "MYVALUE", value)
}

func TestParseConfigurationLineWithSingleQuotes(t *testing.T) {
	key, value := parseConfigurationLine("KEY='MYVALUE'")

	assert.Equal(t, "KEY", key)
	assert.Equal(t, "MYVALUE", value)
}

func TestProcessLineWithComment(t *testing.T) {
	testString := "Test"
	line := "#LOG_LEVEL=" + testString
	processLine(line)

	assert.Equal(t, "", models.GetConfiguration().LogLevel)
}

func TestProcessLineWithValidLine(t *testing.T) {
	testString := "Test"
	line := "LOG_LEVEL=" + testString
	processLine(line)

	assert.Equal(t, testString, models.GetConfiguration().LogLevel)
}

func TestProcessLineWithValidLineAndSpaces(t *testing.T) {
	testStringWithoutSpace := "Test"
	testString := " " + testStringWithoutSpace
	line := "LOG_LEVEL =" + testString
	processLine(line)

	assert.Equal(t, testStringWithoutSpace, models.GetConfiguration().LogLevel)
}
