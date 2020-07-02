package config

import "testing"

func TestParseIpConfigurationLine(t *testing.T) {
	var namespaceToCheck = "mynamespace"
	var ipToCheck = "1.2.3.4"
	var line = namespaceToCheck + " " + ipToCheck
	namespace, ip := parseIpConfigurationLine(line)

	if namespace != namespaceToCheck {
		t.Errorf("Failed. Namespace is [%s] instead of [%s]", namespace, namespaceToCheck)
	} else if ip != ipToCheck {
		t.Errorf("Failed. Ip is [%s] instead of [%s]", ip, ipToCheck)
	} else {
		t.Logf("Success. Found namespace [%s] and IP [%s]", namespace, ip)
	}
}

func TestParseIpConfigurationWithEqual(t *testing.T) {
	var namespaceToCheck = "mynamespace"
	var ipToCheck = "1.2.3.4"
	var line = namespaceToCheck + "=" + ipToCheck
	namespace, ip := parseIpConfigurationLine(line)

	if namespace == namespaceToCheck && ip == ipToCheck {
		t.Logf("Success. Namespace is [%s] and IP is [%s].", namespaceToCheck, ip)
	} else {
		t.Errorf("Failed. Namespace is [%s] and IP is [%s], but both should be empty!", namespace, ip)
	}
}

func TestParseIpConfigurationWithSpaces(t *testing.T) {
	var namespaceToCheck = "mynamespace"
	var ipToCheck = "1.2.3.4"
	var line = namespaceToCheck + " = " + ipToCheck
	namespace, ip := parseIpConfigurationLine(line)

	if namespace == namespaceToCheck && ip == ipToCheck {
		t.Logf("Success. Namespace is [%s] and IP is [%s].", namespaceToCheck, ip)
	} else {
		t.Errorf("Failed. Namespace is [%s] and IP is [%s], but both should be empty!", namespace, ip)
	}
}

func TestParseIpConfigurationWithInvalidLine(t *testing.T) {
	var namespaceToCheck = "mynamespace"
	var line = namespaceToCheck
	namespace, ip := parseIpConfigurationLine(line)

	if namespace == "" && ip == "" {
		t.Log("Success. Both values are empty.")
	} else {
		t.Errorf("Failed. Namespace is [%s] and IP is [%s], but both should be empty!", namespace, ip)
	}
}
