package installactions

import (
	"k8s-management-go/app/constants"
	"k8s-management-go/app/utils/scripts"
)

// ActionShellScriptsInstall is the action to execute the install shell scripts
func ActionShellScriptsInstall(namespace string) (err error) {
	return scripts.ExecuteScriptsInstallScriptsForNamespace(namespace, constants.DirProjectScriptsInstallPrefix)
}
