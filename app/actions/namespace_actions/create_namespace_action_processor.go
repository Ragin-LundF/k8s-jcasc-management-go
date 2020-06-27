package namespace_actions

import (
	"k8s-management-go/app/actions/install_actions"
	"k8s-management-go/app/models"
	"k8s-management-go/app/utils/loggingstate"
)

func ProcessNamespaceCreation(state models.StateData) (err error) {
	loggingstate.AddInfoEntry("Start creating namespace...")

	loggingstate.AddInfoEntry("-> Check and create namespace if necessary...")
	if err = install_actions.ActionCheckAndCreateNamespace(state.Namespace); err != nil {
		loggingstate.AddErrorEntryAndDetails("-> Check and create namespace if necessary...failed", err.Error())
		return err
	}
	loggingstate.AddInfoEntry("-> Check and create namespace if necessary...done")
	loggingstate.AddInfoEntry("Start creating namespace...done")
	return nil
}
