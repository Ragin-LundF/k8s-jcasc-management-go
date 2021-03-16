package install

import (
	"k8s-management-go/app/constants"
	"k8s-management-go/app/utils/scripts"
)

// ActionShellScriptsInstall is the action to execute the install shell scripts
func (projectConfig *ProjectConfig) ActionShellScriptsInstall() (err error) {
	return scripts.ExecuteScriptsInstallScriptsForNamespace(
		projectConfig.Project.Base.Namespace,
		constants.DirProjectScriptsInstallPrefix)
}
