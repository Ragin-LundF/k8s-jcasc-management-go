package createproject

import (
	"k8s-management-go/app/cli/dialogs"
	"k8s-management-go/app/utils/logger"
	"k8s-management-go/app/utils/loggingstate"
	"k8s-management-go/app/utils/validator"
)

func IpAddressWorkflow() (ipAddress string, err error) {
	log := logger.Log()
	// Validator for IP address
	validate := validator.ValidateIp

	// Prepare prompt
	dialogs.ClearScreen()
	ipAddress, err = dialogs.DialogPrompt("Enter the load balancer IP address", validate)
	// check if everything was ok
	if err != nil {
		loggingstate.AddErrorEntryAndDetails("  -> Unable to get the IP address.", err.Error())
		log.Errorf("[IpAddressWorkflow] Unable to get the IP address. %s\n", err.Error())
		return ipAddress, err
	}

	return ipAddress, nil
}
