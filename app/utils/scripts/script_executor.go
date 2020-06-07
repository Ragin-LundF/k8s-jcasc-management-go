package scripts

import (
	"k8s-management-go/app/cli/loggingstate"
	"k8s-management-go/app/constants"
	"k8s-management-go/app/models"
	"k8s-management-go/app/utils/files"
	"k8s-management-go/app/utils/logger"
	"os/exec"
)

func ExecuteScriptsInstallScriptsForNamespace(namespace string, filePrefix string) (err error) {
	log := logger.Log()
	log.Info("[Execute Scripts] Try to execute scripts for namespace [%v]...", namespace)
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
		log.Info("[Execute Scripts] Script directory is available for namespace [%v]...", namespace)
		// prepare file filter for install
		scriptFileEnding := constants.ScriptsFileEnding
		var fileFilter = files.FileFilter{
			Prefix: &filePrefix,
			Suffix: &scriptFileEnding,
		}
		// list files which match to filter
		fileArray, err := files.ListFilesOfDirectoryWithFilter(scriptFolder, &fileFilter)
		if err != nil {
			log.Error("[Execute Scripts] Error while filtering for files with prefix [%v] in directory [%v].", filePrefix, scriptFolder)
			return err
		}

		// if array has content -> execute scripts
		if cap(*fileArray) > 0 {
			// iterate over filtered file array and execute scripts
			for _, file := range *fileArray {
				scriptWithPath := files.AppendPath(scriptFolder, file)
				log.Info("[Execute Scripts] Trying to execute script [%v]", scriptWithPath)
				loggingstate.AddInfoEntryAndDetails("  -> Try to execute script ["+file+"]...", scriptWithPath)

				// Execute scripts
				_ = exec.Command("chmod", "755", scriptWithPath).Run()
				outputCmd, err := exec.Command(scriptWithPath).CombinedOutput()
				if err != nil {
					loggingstate.AddErrorEntryAndDetails("  -> Try to execute script ["+file+"]...failed. See output.", string(outputCmd))
					loggingstate.AddErrorEntryAndDetails("  -> Try to execute script ["+file+"]...failed. See error.", err.Error())
					log.Error("Unable to execute script [%v] Error: \n%v", scriptWithPath, err)
					log.Error("Unable to execute script [%v] Output: \n%v", scriptWithPath, string(outputCmd))

					return err
				}
				loggingstate.AddInfoEntryAndDetails("  -> Try to execute script ["+file+"]...done. See output.", string(outputCmd))
				log.Info("[Execute Scripts] Script output of [%v]: \n%v", scriptWithPath, outputCmd)
			}
		} else {
			loggingstate.AddInfoEntry("  -> No scripts with prefix [" + filePrefix + "] found for [" + namespace + "]")
		}
	} else {
		loggingstate.AddInfoEntry("  -> No scripts directory found for [" + namespace + "]")
	}

	log.Info("[Execute Scripts] Executing scripts for namespace [%v] done.", namespace)
	return nil
}
