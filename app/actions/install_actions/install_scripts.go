package install_actions

import (
	"k8s-management-go/app/constants"
	"k8s-management-go/app/utils/scripts"
)

func ShellScriptsInstall(namespace string) (err error) {
	return scripts.ExecuteScriptsInstallScriptsForNamespace(namespace, constants.DirProjectScriptsInstallPrefix)
}
