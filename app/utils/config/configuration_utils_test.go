package config

import (
	"k8s-management-go/app/models"
	"testing"
)

// Read configuration from k8s-management config file
func TestParseConfigurationLine(t *testing.T) {
	key, value := parseConfigurationLine("KEY=MYVALUE")

	if key != "KEY" && value != "MYVALUE" {
		t.Errorf("Unable to parse configuration. Key [%v] Value[%v]", key, value)
	} else {
		t.Log("Success. Config was parsed and split into correct key/value pair.")
	}
}

func TestParseConfigurationLineWithDoubleQuotes(t *testing.T) {
	key, value := parseConfigurationLine("KEY=\"MYVALUE\"")

	if key != "KEY" && value != "MYVALUE" {
		t.Errorf("Unable to parse configuration. Key [%v] Value[%v]", key, value)
	} else {
		t.Log("Success. Config was parsed and split into correct key/value pair.")
	}
}

func TestParseConfigurationLineWithSingleQuotes(t *testing.T) {
	key, value := parseConfigurationLine("KEY='MYVALUE'")

	if key != "KEY" && value != "MYVALUE" {
		t.Errorf("Unable to parse configuration. Key [%v] Value[%v]", key, value)
	} else {
		t.Log("Success. Config was parsed and split into correct key/value pair.")
	}
}

func TestProcessLineWithComment(t *testing.T) {
	testString := "Test"
	line := "#LOG_LEVEL=" + testString
	processLine(line)
	if models.GetConfiguration().LogLevel != "" {
		t.Errorf("Failed. LogLevel was set to [%s]", models.GetConfiguration().LogLevel)
	} else {
		t.Log("Success. No LogLevel was set.")
	}
}

func TestProcessLineWithValidLine(t *testing.T) {
	testString := "Test"
	line := "LOG_LEVEL=" + testString
	processLine(line)
	if models.GetConfiguration().LogLevel != testString {
		t.Errorf("Failed. LogLevel was set to [%s]", models.GetConfiguration().LogLevel)
	} else {
		t.Logf("Success. No LogLevel should be [%s] and was set to [%s].", testString, models.GetConfiguration().LogLevel)
	}
}

func TestProcessLineWithValidLineAndSpaces(t *testing.T) {
	testStringWithoutSpace := "Test"
	testString := " " + testStringWithoutSpace
	line := "LOG_LEVEL =" + testString
	processLine(line)
	if models.GetConfiguration().LogLevel != testStringWithoutSpace {
		t.Errorf("Failed. LogLevel was set to [%s]", models.GetConfiguration().LogLevel)
	} else {
		t.Logf("Success. No LogLevel should be [%s] and was set to [%s].", testStringWithoutSpace, models.GetConfiguration().LogLevel)
	}
}
