package install_actions

import (
	"k8s-management-go/app/constants"
	"k8s-management-go/app/utils/scripts"
)

func ActionShellScriptsInstall(namespace string) (err error) {
	return scripts.ExecuteScriptsInstallScriptsForNamespace(namespace, constants.DirProjectScriptsInstallPrefix)
}
