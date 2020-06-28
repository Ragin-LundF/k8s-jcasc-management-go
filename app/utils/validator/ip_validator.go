package validator

import (
	"errors"
	"k8s-management-go/app/models"
	"regexp"
	"strings"
)

func ValidateIp(input string) error {
	// check if IP address has correct format
	regex := regexp.MustCompile(`^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$`)
	if !regex.Match([]byte(input)) {
		return errors.New("IP address is not valid! ")
	}
	// check, that ip address was not already used
	for _, ipConfig := range models.GetIpConfiguration().Ips {
		if strings.ToLower(ipConfig.Ip) == strings.ToLower(input) {
			return errors.New("IP address already in use! ")
		}
	}
	return nil
}
