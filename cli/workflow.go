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
func Workflow(info string, err error) {
	selectedCommand := Menu(info, err)
	info, err = startCommandAction(selectedCommand)
	// recall Workflow to show menu after finished actions
	Workflow(info, err)
}

// process the selected command and start the action
func startCommandAction(command string) (info string, err error) {
	// evaluate the command
	switch command {
	case constants.CommandInstall:
		info, err = InstallJenkins()
	case constants.CommandUninstall:
		fmt.Println("start uninstall")
	case constants.CommandUpgrade:
		fmt.Println("start upgrade")
	case constants.CommandEncryptSecrets:
		info, err = EncryptSecretsFile()
	case constants.CommandDecryptSecrets:
		info, err = DecryptSecretsFile()
	case constants.CommandApplySecrets:
		info, err = ApplySecrets()
	case constants.CommandApplySecretsToAll:
		info, err = ApplySecretsToAllNamespaces()
	case constants.CommandCreateProject:
		fmt.Println("start create project")
	case constants.CommandCreateDeploymentOnlyProject:
		fmt.Println("start create deployment project")
	case constants.CommandCreateJenkinsUserPassword:
		info, err = CreateJenkinsUserPassword()
	case constants.CommandQuit:
		os.Exit(0)
	case constants.ErrorPromptFailed:
		os.Exit(0)
	}
	return info, err
}
