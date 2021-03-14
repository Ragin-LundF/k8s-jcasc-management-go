package install

import (
	"k8s-management-go/app/constants"
	"k8s-management-go/app/utils/scripts"
)

// ActionShellScriptsUninstall executes the actions to execute the uninstall scripts
func (projectConfig *ProjectConfig) ActionShellScriptsUninstall() (err error) {
	return scripts.ExecuteScriptsInstallScriptsForNamespace(
		projectConfig.Project.Base.Namespace,
		constants.DirProjectScriptsUninstallPrefix)
}
