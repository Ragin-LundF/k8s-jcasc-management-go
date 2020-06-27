package namespace

import (
	"k8s-management-go/app/actions/namespace_actions"
	"k8s-management-go/app/cli/createproject"
	"k8s-management-go/app/models"
	"k8s-management-go/app/utils/loggingstate"
)

func WorkflowCreateNamespace() (err error) {
	var state models.StateData
	state.Namespace, err = createproject.ProjectWizardAskForNamespace()
	if err != nil {
		loggingstate.AddErrorEntryAndDetails("-> AskForNamespace dialog aborted...", err.Error())
		loggingstate.LogLoggingStateEntries()
	}

	err = namespace_actions.ProcessNamespaceCreation(state)
	loggingstate.LogLoggingStateEntries()

	return nil
}
