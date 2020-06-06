package createproject

import (
	"k8s-management-go/app/constants"
	"k8s-management-go/app/utils/logger"
)

func ProjectWizardWorkflow(deploymentOnly bool) (info string, err error) {
	log := logger.Log()

	// Ask for namespace
	info = logger.InfoLog(info, "[ProjectWizardWorkflow] Starting Project Wizard.")
	info = logger.InfoLog(info, "[ProjectWizardWorkflow] -> Ask for namespace...")
	namespace, err := ProjectWizardAskForNamespace()
	if err != nil {
		log.Error("[ProjectWizardWorkflow] Project wizard aborted because of errors while getting namespace name.")
		info = logger.InfoLog(info, "[ProjectWizardWorkflow] -> Ask for namespace...See errors")
		return info, err
	}
	info = logger.InfoLog(info, "[ProjectWizardWorkflow] -> Ask for namespace...done")

	// Ask for IP address
	info = logger.InfoLog(info, "[ProjectWizardWorkflow] -> Ask for IP address...")
	ipAddress, err := ProjectWizardAskForIpAddress()
	if err != nil {
		log.Error("[ProjectWizardWorkflow] Project wizard aborted because of errors while getting IP address.")
		info = logger.InfoLog(info, "[ProjectWizardWorkflow] -> Ask for IP address...See errors")
		return info, err
	}
	info = logger.InfoLog(info, "[ProjectWizardWorkflow] -> Ask for IP address...done")

	// declare vars for next if statement
	var infoLog string
	var jenkinsSystemMsg string
	var jobsCfgRepo string
	var selectedCloudTemplates []string
	var existingPvc string

	// if it is not only a deployment project ask for other Jenkins related vars
	if !deploymentOnly {
		// Select cloud templates
		info = logger.InfoLog(info, "[ProjectWizardWorkflow] -> Ask for cloud templates...")
		selectedCloudTemplates, infoLog, err = ProjectWizardAskForCloudTemplates()
		info = info + constants.NewLine + infoLog
		if err != nil {
			log.Error("[ProjectWizardWorkflow] Project wizard aborted because of errors while getting cloud templates.")
			info = logger.InfoLog(info, "[ProjectWizardWorkflow] -> Ask for cloud templates...See errors")
			return info, err
		}
		info = logger.InfoLog(info, "[ProjectWizardWorkflow] -> Ask for cloud templates...done")

		// Ask for existing persistent volume claim (PVC)
		info = logger.InfoLog(info, "[ProjectWizardWorkflow] -> Ask for persistent volume claim...")
		existingPvc, err = ProjectWizardAskForExistingPersistentVolumeClaim()
		if err != nil {
			log.Error("[ProjectWizardWorkflow] Project wizard aborted because of errors while getting persistent volume claim.")
			info = logger.InfoLog(info, "[ProjectWizardWorkflow] -> Ask for persistent volume claim...See errors")
			return info, err
		}
		info = logger.InfoLog(info, "[ProjectWizardWorkflow] -> Ask for persistent volume claim...done")

		// Ask for Jenkins system message
		info = logger.InfoLog(info, "[ProjectWizardWorkflow] -> Ask for Jenkins system message...")
		jenkinsSystemMsg, err = ProjectWizardAskForJenkinsSystemMessage(namespace)
		if err != nil {
			log.Error("[ProjectWizardWorkflow] Project wizard aborted because of errors while getting Jenkins system message.")
			info = logger.InfoLog(info, "[ProjectWizardWorkflow] -> Ask for Jenkins system message...See errors")
			return info, err
		}
		info = logger.InfoLog(info, "[ProjectWizardWorkflow] -> Ask for Jenkins system message...done")

		// Ask for Jobs Configuration repository
		info = logger.InfoLog(info, "[ProjectWizardWorkflow] -> Ask for jobs configuration repository...")
		jobsCfgRepo, err = ProjectWizardAskForJobsConfigurationRepository()
		if err != nil {
			log.Error("[ProjectWizardWorkflow] Project wizard aborted because of errors while getting jobs configuration repository.")
			info = logger.InfoLog(info, "[ProjectWizardWorkflow] -> Ask for jobs configuration repository...See errors")
			return info, err
		}
		info = logger.InfoLog(info, "[ProjectWizardWorkflow] -> Ask for jobs configuration repository...done")
	}

	// Process data and create project
	infoLog, err = ProcessProjectCreate(namespace, ipAddress, jenkinsSystemMsg, jobsCfgRepo, existingPvc, selectedCloudTemplates, deploymentOnly)
	info = info + constants.NewLine + infoLog

	return info, err
}
