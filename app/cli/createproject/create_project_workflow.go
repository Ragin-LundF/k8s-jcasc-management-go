package createproject

import (
	"k8s-management-go/app/constants"
	"k8s-management-go/app/utils/logger"
)

func ProjectWizardWorkflow() (info string, err error) {
	log := logger.Log()

	// Ask for namespace
	info = info + constants.NewLine + "[ProjectWizardWorkflow] Starting Project Wizard."
	info = info + constants.NewLine + "[ProjectWizardWorkflow] -> Ask for namespace..."
	namespace, err := ProjectWizardAskForNamespace()
	if err != nil {
		log.Error("[ProjectWizardWorkflow] Project wizard aborted because of errors while getting namespace name.")
		info = info + constants.NewLine + "[ProjectWizardWorkflow] -> Ask for namespace...error!"
		return info, err
	}
	info = info + constants.NewLine + "[ProjectWizardWorkflow] -> Ask for namespace...done"

	// Ask for IP address
	info = info + constants.NewLine + "[ProjectWizardWorkflow] -> Ask for IP address..."
	ipAddress, err := ProjectWizardAskForIpAddress()
	if err != nil {
		log.Error("[ProjectWizardWorkflow] Project wizard aborted because of errors while getting IP address.")
		info = info + constants.NewLine + "[ProjectWizardWorkflow] -> Ask for IP address...error!"
		return info, err
	}
	info = info + constants.NewLine + "[ProjectWizardWorkflow] -> Ask for IP address...done"

	// Ask for Jenkins system message
	info = info + constants.NewLine + "[ProjectWizardWorkflow] -> Ask for Jenkins system message..."
	jenkinsSystemMsg, err := ProjectWizardAskForJenkinsSystemMessage(namespace)
	if err != nil {
		log.Error("[ProjectWizardWorkflow] Project wizard aborted because of errors while getting Jenkins system message.")
		info = info + constants.NewLine + "[ProjectWizardWorkflow] -> Ask for Jenkins system message...error!"
		return info, err
	}
	info = info + constants.NewLine + "[ProjectWizardWorkflow] -> Ask for Jenkins system message...done"

	// Ask for Jobs Configuration repository
	info = info + constants.NewLine + "[ProjectWizardWorkflow] -> Ask for jobs configuration repository..."
	jobsCfgRepo, err := ProjectWizardAskForJobsConfigurationRepository()
	if err != nil {
		log.Error("[ProjectWizardWorkflow] Project wizard aborted because of errors while getting jobs configuration repository.")
		info = info + constants.NewLine + "[ProjectWizardWorkflow] -> Ask for jobs configuration repository...error!"
		return info, err
	}
	info = info + constants.NewLine + "[ProjectWizardWorkflow] -> Ask for jobs configuration repository...done"

	// Ask for existing persistent volume claim (PVC)
	info = info + constants.NewLine + "[ProjectWizardWorkflow] -> Ask for persistent volume claim..."
	existingPvc, err := ProjectWizardAskForExistingPersistentVolumeClaim()
	if err != nil {
		log.Error("[ProjectWizardWorkflow] Project wizard aborted because of errors while getting persistent volume claim.")
		info = info + constants.NewLine + "[ProjectWizardWorkflow] -> Ask for persistent volume claim...error!"
		return info, err
	}
	info = info + constants.NewLine + "[ProjectWizardWorkflow] -> Ask for persistent volume claim...done"

	// Select cloud templates
	info = info + constants.NewLine + "[ProjectWizardWorkflow] -> Ask for cloud templates..."
	selectedCloudTemplates, infoLog, err := ProjectWizardAskForCloudTemplates()
	info = info + constants.NewLine + infoLog
	if err != nil {
		log.Error("[ProjectWizardWorkflow] Project wizard aborted because of errors while getting cloud templates.")
		info = info + constants.NewLine + "[ProjectWizardWorkflow] -> Ask for cloud templates...error!"
		return info, err
	}
	info = info + constants.NewLine + "[ProjectWizardWorkflow] -> Ask for cloud templates...done"

	// Process data and create project
	infoLog, err = ProcessProjectCreate(namespace, ipAddress, jenkinsSystemMsg, jobsCfgRepo, existingPvc, selectedCloudTemplates)
	info = info + constants.NewLine + infoLog

	return info, err
}
