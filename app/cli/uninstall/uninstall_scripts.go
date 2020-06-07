package uninstall

import (
	"k8s-management-go/app/constants"
	"k8s-management-go/app/utils/scripts"
)

func ShellScriptsUninstall(namespace string) (err error) {
	return scripts.ExecuteScriptsInstallScriptsForNamespace(namespace, constants.DirProjectScriptsUninstallPrefix)
}
