package createproject

import (
	"k8s-management-go/app/actions/project"
	"k8s-management-go/app/cli/dialogs"
	"k8s-management-go/app/constants"
)

// CloudTemplatesWorkflow is the project wizard dialog for cloud templates
func CloudTemplatesWorkflow() (cloudTemplates []string, err error) {
	// look if cloud templates are available
	var cloudTemplateDialog dialogs.CloudTemplatesDialog
	var cloudTemplateFileList = project.ActionReadCloudTemplates()

	// prepare selection
	cloudTemplateSelectionList := []string{constants.ActionSelectTemplatesBelowOrContinue}

	// if cloud templates were found add them to the new select list and ask which should be used
	if len(cloudTemplateSelectionList) > 0 {
		cloudTemplateSelectionList = append(cloudTemplateSelectionList, cloudTemplateFileList...)
		cloudTemplateDialog.CloudTemplateFiles = cloudTemplateSelectionList

		_ = dialogs.DialogAskForCloudTemplates(&cloudTemplateDialog)
	}

	return cloudTemplateDialog.SelectedCloudTemplates, nil
}
