package install

import (
	"errors"
	"fmt"
	"k8s-management-go/app/actions/project"
	"k8s-management-go/app/configuration"
	"k8s-management-go/app/constants"
	"k8s-management-go/app/utils/helm"
	"k8s-management-go/app/utils/loggingstate"
)

// ActionHelmInstallJenkins installs Jenkins with Helm
func (projectConfig *ProjectConfig) ActionHelmInstallJenkins() (err error) {
	loggingstate.AddInfoEntry(fmt.Sprintf(
		"-> Try to %s Jenkins on namespace [%s] with deployment name [%s]...",
		projectConfig.HelmCommand,
		projectConfig.Project.Base.Namespace,
		projectConfig.Project.Base.DeploymentName))

	// check if command is ok
	if projectConfig.HelmCommand == constants.HelmCommandInstall || projectConfig.HelmCommand == constants.HelmCommandUpgrade {
		// prepare files and directories
		var helmChartsJenkinsDirectory = configuration.GetConfiguration().FilePathWithBasePath(constants.DirHelmJenkinsMaster)
		// prepare file directories
		var helmChartsJenkinsValuesFile string
		helmChartsJenkinsValuesFile, err = projectConfig.PrepareInstallYAML(constants.FilenameJenkinsHelmValues)
		if err != nil {
			project.RemoveTempFile(helmChartsJenkinsValuesFile)
			return err
		}

		// execute Helm command
		var argsForCommand = []string{
			projectConfig.Project.Base.DeploymentName,
			helmChartsJenkinsDirectory,
		}
		if configuration.GetConfiguration().K8SManagement.DryRunOnly {
			argsForCommand = append(
				argsForCommand,
				"--dry-run",
				"--debug")
		}
		argsForCommand = append(
			argsForCommand,
			"-n",
			projectConfig.Project.Base.Namespace,
			"-f",
			helmChartsJenkinsValuesFile)

		// executing jenkins helm install
		loggingstate.AddInfoEntry(fmt.Sprintf(
			"-> Start installing/upgrading Jenkins with Helm on namespace [%s]...",
			projectConfig.Project.Base.Namespace))
		err = helm.ExecutorHelm(projectConfig.HelmCommand, argsForCommand)

		project.RemoveTempFile(helmChartsJenkinsValuesFile)
		if err != nil {
			loggingstate.AddErrorEntryAndDetails(fmt.Sprintf(
				"-> Unable to install/upgrade Jenkins on namespace [%s] with deployment name [%s]",
				projectConfig.Project.Base.Namespace,
				projectConfig.Project.Base.DeploymentName), err.Error())
			return err
		}

		loggingstate.AddInfoEntry(fmt.Sprintf(
			"-> Start installing/upgrading Jenkins with Helm on namespace [%s]...done",
			projectConfig.Project.Base.Namespace))
	} else {
		// helm command was wrong -> abort
		loggingstate.AddErrorEntry(fmt.Sprintf(
			"-> Try to install/upgrade Jenkins on namespace [%s] with deployment name [%s]...Wrong command [%s]",
			projectConfig.Project.Base.Namespace,
			projectConfig.Project.Base.DeploymentName,
			projectConfig.HelmCommand))
		return errors.New("Helm command [" + projectConfig.HelmCommand + "] unknown.")
	}
	loggingstate.AddInfoEntry(fmt.Sprintf(
		"-> Try to %s Jenkins on namespace [%s] with deployment name [%s]...done",
		projectConfig.HelmCommand,
		projectConfig.Project.Base.Namespace,
		projectConfig.Project.Base.DeploymentName))

	return err
}
