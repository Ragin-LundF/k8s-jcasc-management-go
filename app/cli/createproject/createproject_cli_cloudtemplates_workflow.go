package createproject

import (
	"k8s-management-go/app/actions/createprojectactions"
	"k8s-management-go/app/cli/dialogs"
)

// CloudTemplatesWorkflow is the project wizard dialog for cloud templates
func CloudTemplatesWorkflow() (cloudTemplates []string, err error) {
	// look if cloud templates are available
	var cloudTemplateDialog dialogs.CloudTemplatesDialog
	cloudTemplateFileList := createprojectactions.ActionReadCloudTemplates()

	// prepare selection
	cloudTemplateSelectionList := []string{"===== Select templates below or continue here ======"}

	// if cloud templates were found add them to the new select list and ask which should be used
	if cap(cloudTemplateSelectionList) > 0 {
		cloudTemplateSelectionList = append(cloudTemplateSelectionList, cloudTemplateFileList...)
		cloudTemplateDialog.CloudTemplateFiles = cloudTemplateSelectionList

		_ = dialogs.DialogAskForCloudTemplates(&cloudTemplateDialog)
	}

	return cloudTemplateDialog.SelectedCloudTemplates, nil
}
