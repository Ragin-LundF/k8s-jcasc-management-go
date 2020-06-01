package cli

import (
	"fmt"
	"k8s-management-go/constants"
	"os"
)

type state struct {
	Command  string
	Previous string
}

// Workflow entrypoint
func Workflow() {
	selectedCommand := Menu()
	startCommandAction(selectedCommand)
	// recall Workflow to show menu after finished actions
	Workflow()
}

// process the selected command and start the action
func startCommandAction(command string) {
	switch command {
	case constants.CommandInstall:
		fmt.Println("start install")
	case constants.CommandUninstall:
		fmt.Println("start uninstall")
	case constants.CommandUpgrade:
		fmt.Println("start uninstall")
	case constants.CommandEncryptSecrets:
		fmt.Println("start uninstall")
	case constants.CommandDecryptSecrets:
		fmt.Println("start uninstall")
	case constants.CommandApplySecrets:
		fmt.Println("start uninstall")
	case constants.CommandApplySecretsToAll:
		fmt.Println("start uninstall")
	case constants.CommandCreateProject:
		fmt.Println("start uninstall")
	case constants.CommandCreateDeploymentOnlyProject:
		fmt.Println("start uninstall")
	case constants.CommandCreateJenkinsUserPassword:
		CreateJenkinsUserPassword()
	case constants.CommandQuit:
		os.Exit(0)
	case constants.ErrorPromptFailed:
		os.Exit(0)
	}
}
