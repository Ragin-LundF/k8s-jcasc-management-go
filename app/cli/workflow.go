package cli

import (
	"k8s-management-go/app/cli/createproject"
	"k8s-management-go/app/cli/dialogs"
	"k8s-management-go/app/cli/install"
	"k8s-management-go/app/cli/jenkinsuser"
	"k8s-management-go/app/cli/loggingstate"
	"k8s-management-go/app/cli/menu"
	"k8s-management-go/app/cli/namespace"
	"k8s-management-go/app/cli/secrets"
	"k8s-management-go/app/cli/uninstall"
	"k8s-management-go/app/constants"
	"os"
)

// Workflow entrypoint
func Workflow(info string, err error) {
	selectedCommand := menu.Menu(info, err)
	err = startCommandAction(selectedCommand)
	// show output
	dialogs.DialogShowLogging(loggingstate.GetLoggingStateEntries(), err)
	loggingstate.ClearLoggingState()
	// recall Workflow to show menu after finished actions
	Workflow(info, err)
}

// process the selected command and start the action
func startCommandAction(command string) (err error) {
	// evaluate the command
	switch command {
	case constants.CommandInstall:
		err = install.DoUpgradeOrInstall(constants.HelmCommandInstall)
	case constants.CommandUninstall:
		err = uninstall.DoUninstall()
	case constants.CommandUpgrade:
		err = install.DoUpgradeOrInstall(constants.HelmCommandUpgrade)
	case constants.CommandEncryptSecrets:
		err = secrets.EncryptSecretsFile()
	case constants.CommandDecryptSecrets:
		err = secrets.DecryptSecretsFile()
	case constants.CommandApplySecrets:
		err = secrets.ApplySecrets()
	case constants.CommandApplySecretsToAll:
		err = secrets.ApplySecretsToAllNamespaces()
	case constants.CommandCreateNamespace:
		err = namespace.WorkflowCreateNamespace()
	case constants.CommandCreateProject:
		err = createproject.ProjectWizardWorkflow(false)
	case constants.CommandCreateDeploymentOnlyProject:
		err = createproject.ProjectWizardWorkflow(true)
	case constants.CommandCreateJenkinsUserPassword:
		err = jenkinsuser.CreateJenkinsUserPassword()
	case constants.CommandQuit:
		os.Exit(0)
	case constants.ErrorPromptFailed:
		os.Exit(0)
	}
	return err
}
