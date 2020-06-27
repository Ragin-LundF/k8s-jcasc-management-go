package namespace

import (
	"k8s-management-go/app/actions/install_actions"
	"k8s-management-go/app/cli/createproject"
	"k8s-management-go/app/utils/loggingstate"
)

func WorkflowCreateNamespace() (err error) {
	loggingstate.AddInfoEntry("Start creating namespace...")
	namespace, err := createproject.ProjectWizardAskForNamespace()
	if err != nil {
		loggingstate.AddErrorEntryAndDetails("-> AskForNamespace dialog aborted...", err.Error())
	}

	loggingstate.AddInfoEntry("-> Check and create namespace if necessary...")
	if err = install_actions.CheckAndCreateNamespace(namespace); err != nil {
		loggingstate.AddErrorEntryAndDetails("-> Check and create namespace if necessary...failed", err.Error())
		return err
	}
	loggingstate.AddInfoEntry("-> Check and create namespace if necessary...done")
	loggingstate.AddInfoEntry("Start creating namespace...done")
	return nil
}
