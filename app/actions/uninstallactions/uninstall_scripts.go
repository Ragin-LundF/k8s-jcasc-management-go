package uninstallactions

import (
	"k8s-management-go/app/constants"
	"k8s-management-go/app/utils/scripts"
)

// ActionShellScriptsUninstall executes the actions to execute the uninstall scripts
func ActionShellScriptsUninstall(namespace string) (err error) {
	return scripts.ExecuteScriptsInstallScriptsForNamespace(namespace, constants.DirProjectScriptsUninstallPrefix)
}
