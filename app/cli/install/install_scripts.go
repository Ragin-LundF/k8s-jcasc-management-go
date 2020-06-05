package install

import (
	"errors"
	"k8s-management-go/app/constants"
	"k8s-management-go/app/models"
	"k8s-management-go/app/utils/files"
	"k8s-management-go/app/utils/logger"
	"os/exec"
)

func ShellScriptsInstall(namespace string) (info string, err error) {
	log := logger.Log()
	log.Info("[Install Scripts] Searching for install scripts for namespace [" + namespace + "] ")
	// calculate path to script folder
	var scriptFolder = files.AppendPath(
		files.AppendPath(
			models.GetProjectBaseDirectory(),
			namespace,
		),
		constants.DirProjectScripts,
	)

	// check if folder exists
	log.Info("[Install Scripts] Search for shell scripts in folder [" + scriptFolder + "]")
	var isScriptsDirectoryAvailable = files.FileOrDirectoryExists(scriptFolder)
	if isScriptsDirectoryAvailable {
		log.Info("[Install Scripts] Found folder [" + scriptFolder + "]. Start searching for files...")
		// prepare file filter for install
		filePrefix := constants.DirProjectScriptsInstallPrefix
		scriptFileEnding := constants.ScriptsFileEnding
		var fileFilter = files.FileFilter{
			Prefix: &filePrefix,
			Suffix: &scriptFileEnding,
		}
		// list files which match to filter
		fileArray, err := files.ListFilesOfDirectoryWithFilter(scriptFolder, &fileFilter)
		if err != nil {
			log.Error("[Install Scripts] Error while filtering for install files in directory [" + scriptFolder + "].")
			return info, err
		}

		// iterate over filtered file array and execute scripts
		for _, file := range *fileArray {
			scriptWithPath := files.AppendPath(scriptFolder, file)
			log.Info("[Install Scripts] Trying to execute install script [" + scriptWithPath + "]")
			info = info + constants.NewLine + "Trying to execute install script [" + scriptWithPath + "]"
			// Execute scripts
			outputCmd, err := exec.Command("sh", "-c", scriptWithPath).CombinedOutput()
			info = info + constants.NewLine + string(outputCmd)
			if err != nil {
				err = errors.New(string(outputCmd) + constants.NewLine + err.Error())
				log.Error("Unable to execute script ["+scriptWithPath+"] %v\n", err)

				return info, err
			}
			log.Info("[Install Scripts] Install script output of [" + scriptWithPath + "]:")
			log.Info(outputCmd)
		}
	}
	log.Info("[Install Scripts] Scripts install for namespace [" + namespace + "] done...")
	return info, err
}
