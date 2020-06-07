package uninstall

import (
	"k8s-management-go/app/cli/loggingstate"
	"k8s-management-go/app/constants"
	"k8s-management-go/app/models"
	"k8s-management-go/app/utils/files"
	"k8s-management-go/app/utils/logger"
	"os/exec"
)

func ShellScriptsUninstall(namespace string) (err error) {
	log := logger.Log()
	log.Info("[Uninstall Scripts] Try to uninstall scripts for namespace [%v]...", namespace)
	// calculate path to script folder
	var scriptFolder = files.AppendPath(
		files.AppendPath(
			models.GetProjectBaseDirectory(),
			namespace,
		),
		constants.DirProjectScripts,
	)

	// check if folder exists
	var isScriptsDirectoryAvailable = files.FileOrDirectoryExists(scriptFolder)
	if isScriptsDirectoryAvailable {
		log.Info("[Uninstall Scripts] Uninstall script directory is available for namespace [%v]...", namespace)
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
			log.Error("[Uninstall Scripts] Error while filtering for uninstall files in directory [%v].", scriptFolder)
			return err
		}

		// if array has content -> execute scripts
		if cap(*fileArray) > 0 {
			// iterate over filtered file array and execute scripts
			for _, file := range *fileArray {
				scriptWithPath := files.AppendPath(scriptFolder, file)
				log.Info("[Uninstall Scripts] Trying to execute uninstall script [%v]", scriptWithPath)
				loggingstate.AddInfoEntryAndDetails("  -> Try to execute script ["+file+"]...", scriptWithPath)

				// Execute scripts
				outputCmd, err := exec.Command("sh", "-c", scriptWithPath).CombinedOutput()
				if err != nil {
					loggingstate.AddErrorEntryAndDetails("  -> Try to execute script ["+file+"]...failed. See output.", string(outputCmd))
					log.Error("Unable to execute script [%v] Error: \n%v", scriptWithPath, err)
					log.Error("Unable to execute script [%v] Output: \n%v", scriptWithPath, string(outputCmd))

					return err
				}
				loggingstate.AddInfoEntryAndDetails("  -> Try to execute script ["+file+"]...done. See output.", string(outputCmd))
				log.Info("[Uninstall Scripts] Uninstall script output of [%v]: \n%v", scriptWithPath, outputCmd)
			}
		} else {
			loggingstate.AddInfoEntry("  -> No uninstall scripts found for [" + namespace + "]")
		}
	} else {
		loggingstate.AddInfoEntry("  -> No scripts directory found for [" + namespace + "]")
	}

	log.Info("[Uninstall Scripts] Uninstall scripts for namespace [%v] done.", namespace)
	return nil
}
