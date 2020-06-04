package uninstall

import (
	"errors"
	"k8s-management-go/app/constants"
	"k8s-management-go/app/models/config"
	"k8s-management-go/app/utils/files"
	"os/exec"
)

func ShellScriptsUninstall(namespace string) (info string, err error) {
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
			return info, err
		}

		// iterate over filtered file array and execute scripts
		for _, file := range *fileArray {
			info = info + constants.NewLine + "Uninstalling script [" + namespace + "/" + constants.DirProjectScripts + "/" + file + "]"
			// Execute scripts
			scriptWithPath := files.AppendPath(scriptFolder, file)
			outputCmd, err := exec.Command("sh", "-c", scriptWithPath).CombinedOutput()
			if err != nil {
				err = errors.New(string(outputCmd) + constants.NewLine + err.Error())
				return info, err
			}

			// collect output
			info = info + constants.NewLine + "Output of script [" + scriptWithPath + "]:"
			info = info + constants.NewLine + "==============="
			info = info + string(outputCmd)
			info = info + constants.NewLine + "==============="
		}
	}
	return info, err
}
