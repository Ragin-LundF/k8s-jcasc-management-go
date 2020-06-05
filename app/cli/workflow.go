package cli

import (
	"fmt"
	"k8s-management-go/app/cli/install"
	"k8s-management-go/app/cli/jenkinsuser"
	"k8s-management-go/app/cli/logoutput"
	"k8s-management-go/app/cli/menu"
	"k8s-management-go/app/cli/secrets"
	"k8s-management-go/app/cli/uninstall"
	"k8s-management-go/app/constants"
	"os"
)

type state struct {
	Command  string
	Previous string
}

// Workflow entrypoint
func Workflow(info string, err error) {
	selectedCommand := menu.Menu(info, err)
	info, err = startCommandAction(selectedCommand)
	// show output
	logoutput.DialogShowInfoAndError(info, err)
	// recall Workflow to show menu after finished actions
	Workflow(info, err)
}

// process the selected command and start the action
func startCommandAction(command string) (info string, err error) {
	// evaluate the command
	switch command {
	case constants.CommandInstall:
		info, err = install.DoUpgradeOrInstall(constants.HelmCommandInstall)
	case constants.CommandUninstall:
		info, err = uninstall.DoUninstall()
	case constants.CommandUpgrade:
		info, err = install.DoUpgradeOrInstall(constants.HelmCommandUpgrade)
	case constants.CommandEncryptSecrets:
		info, err = secrets.EncryptSecretsFile()
	case constants.CommandDecryptSecrets:
		info, err = secrets.DecryptSecretsFile()
	case constants.CommandApplySecrets:
		info, err = secrets.ApplySecrets()
	case constants.CommandApplySecretsToAll:
		info, err = secrets.ApplySecretsToAllNamespaces()
	case constants.CommandCreateProject:
		fmt.Println("start create project")
	case constants.CommandCreateDeploymentOnlyProject:
		fmt.Println("start create deployment project")
	case constants.CommandCreateJenkinsUserPassword:
		info, err = jenkinsuser.CreateJenkinsUserPassword()
	case constants.CommandQuit:
		os.Exit(0)
	case constants.ErrorPromptFailed:
		os.Exit(0)
	}
	return info, err
}
