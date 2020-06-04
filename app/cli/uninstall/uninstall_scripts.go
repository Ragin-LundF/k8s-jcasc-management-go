package uninstall

import (
	"errors"
	"k8s-management-go/app/constants"
	"k8s-management-go/app/models/config"
	"k8s-management-go/app/utils/files"
	"k8s-management-go/app/utils/logger"
	"os/exec"
)

func ShellScriptsUninstall(namespace string) (info string, err error) {
	log := logger.Log()
	log.Info("[Uninstall Scripts] Try to uninstall scripts for namespace [" + namespace + "]...")
	// calculate path to script folder
	var scriptFolder = files.AppendPath(
		files.AppendPath(
			config.GetProjectBaseDirectory(),
			namespace,
		),
		constants.DirProjectScripts,
	)

	// check if folder exists
	var isScriptsDirectoryAvailable = files.FileOrDirectoryExists(scriptFolder)
	if isScriptsDirectoryAvailable {
		log.Info("[Uninstall Scripts] Uninstall script directory is available for namespace [" + namespace + "]...")
		// prepare file filter for install
		filePrefix := constants.DirProjectScriptsUninstallPrefix
		scriptFileEnding := constants.ScriptsFileEnding
		var fileFilter = files.FileFilter{
			Prefix: &filePrefix,
			Suffix: &scriptFileEnding,
		}
		// list files which match to filter
		fileArray, err := files.ListFilesOfDirectoryWithFilter(scriptFolder, &fileFilter)
		if err != nil {
			log.Error("[Uninstall Scripts] Error while filtering for uninstall files in directory [" + scriptFolder + "].")
			return info, err
		}

		// iterate over filtered file array and execute scripts
		for _, file := range *fileArray {
			scriptWithPath := files.AppendPath(scriptFolder, file)
			log.Info("[Uninstall Scripts] Trying to execute uninstall script [" + scriptWithPath + "]")
			info = info + constants.NewLine + "Trying to execute uninstall script [" + scriptWithPath + "]"
			// Execute scripts
			outputCmd, err := exec.Command("sh", "-c", scriptWithPath).CombinedOutput()
			info = info + constants.NewLine + string(outputCmd)
			if err != nil {
				err = errors.New(string(outputCmd) + constants.NewLine + err.Error())
				log.Error("Unable to execute script [" + scriptWithPath + "]")
				log.Error(err)

				return info, err
			}
			log.Info("[Uninstall Scripts] Uninstall script output of [" + scriptWithPath + "]:")
			log.Info(outputCmd)
		}
	}

	log.Info("[Uninstall Scripts] Uninstall scripts for namespace [" + namespace + "] done.")
	return info, err
}
