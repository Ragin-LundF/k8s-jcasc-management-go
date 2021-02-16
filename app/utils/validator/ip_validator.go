package validator

import (
	"errors"
	"k8s-management-go/app/models"
	"regexp"
	"strings"
)

// ValidateIP validates if it is a correct IPv4
func ValidateIP(input string) error {
	// Allow empty IP to support free IP assignment
	if input != "" {
		// check if IP address has correct format
		regex := regexp.MustCompile(`^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$`)
		regexDomain := regexp.MustCompile(`^(?:[a-z0-9](?:[a-z0-9-]{0,61}[a-z0-9])?\.)+[a-z0-9][a-z0-9-]{0,61}[a-z0-9]$`)

		if !regex.Match([]byte(input)) && !regexDomain.Match([]byte(input)) {
			return errors.New("IP address or domain is not valid! ")
		}
		// check, that ip address was not already used
		for _, ipConfig := range models.GetIPConfiguration().IPs {
			if strings.ToLower(ipConfig.IP) == strings.ToLower(input) {
				return errors.New("IP address or domain already in use! ")
			}
		}
	}
	return nil
}
