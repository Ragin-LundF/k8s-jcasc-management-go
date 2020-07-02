package namespace

import (
	"k8s-management-go/app/actions/namespaceactions"
	"k8s-management-go/app/cli/createproject"
	"k8s-management-go/app/models"
	"k8s-management-go/app/utils/loggingstate"
)

// WorkflowCreateNamespace is the workflow to create a namespace
func WorkflowCreateNamespace() (err error) {
	var state models.StateData
	state.Namespace, err = createproject.NamespaceWorkflow()
	if err != nil {
		loggingstate.AddErrorEntryAndDetails("-> AskForNamespace dialog aborted...", err.Error())
	}

	err = namespaceactions.ProcessNamespaceCreation(state)

	return nil
}
