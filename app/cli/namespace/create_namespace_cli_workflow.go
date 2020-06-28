package namespace

import (
	"k8s-management-go/app/actions/namespace_actions"
	"k8s-management-go/app/cli/createproject"
	"k8s-management-go/app/models"
	"k8s-management-go/app/utils/loggingstate"
)

func WorkflowCreateNamespace() (err error) {
	var state models.StateData
	state.Namespace, err = createproject.NamespaceWorkflow()
	if err != nil {
		loggingstate.AddErrorEntryAndDetails("-> AskForNamespace dialog aborted...", err.Error())
	}

	err = namespace_actions.ProcessNamespaceCreation(state)

	return nil
}
