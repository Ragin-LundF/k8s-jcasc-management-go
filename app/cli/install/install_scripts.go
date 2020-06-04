package install

import (
	"k8s-management-go/app/constants"
	"k8s-management-go/app/models/config"
	"k8s-management-go/app/utils/files"
	"os/exec"
)

func ShellScriptsInstall(namespace string) (info string, err error) {
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
		filePrefix := constants.DirProjectScriptsInstallPrefix
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
			// Execute scripts
			scriptWithPath := files.AppendPath(scriptFolder, file)
			outputCmd, err := exec.Command("sh", "-c", scriptWithPath).Output()
			if err != nil {
				return info, err
			}

			// collect output
			info = info + "\nOutput of script [" + scriptWithPath + "]:"
			info = info + "\n==============="
			info = info + string(outputCmd)
			info = info + "\n==============="
		}
	}
	return info, err
}
