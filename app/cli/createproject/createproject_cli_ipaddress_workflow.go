package createproject

import (
	"k8s-management-go/app/cli/dialogs"
	"k8s-management-go/app/constants"
	"k8s-management-go/app/utils/loggingstate"
	"k8s-management-go/app/utils/validator"
)

// IPAddressWorkflow represents the ip address workflow
func IPAddressWorkflow() (ipAddress string, err error) {
	// Validator for IP address
	validate := validator.ValidateIP

	// Prepare prompt
	dialogs.ClearScreen()
	ipAddress, err = dialogs.DialogPrompt(constants.TextEnterLoadBalancerIPAddress, validate)
	// check if everything was ok
	if err != nil {
		loggingstate.AddErrorEntryAndDetails(constants.LogErrUnableToGetIPAddress, err.Error())
		return ipAddress, err
	}

	return ipAddress, nil
}
