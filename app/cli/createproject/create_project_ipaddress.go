package createproject

import (
	"errors"
	"k8s-management-go/app/cli/dialogs"
	"k8s-management-go/app/models"
	"k8s-management-go/app/utils/logger"
	"regexp"
	"strings"
)

func ProjectWizardAskForIpAddress() (ipAddress string, err error) {
	log := logger.Log()
	// Validator for IP address
	validate := func(input string) error {
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

	// Prepare prompt
	dialogs.ClearScreen()
	ipAddress, err = dialogs.DialogPrompt("Enter the load balancer IP address", validate)
	// check if everything was ok
	if err != nil {
		log.Error("[ProjectWizardAskForIpAddress] Unable to get the IP address. %v\n", err)
	}

	return ipAddress, err
}
