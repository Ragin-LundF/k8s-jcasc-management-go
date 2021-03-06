package scripts

import (
	"fmt"
	"k8s-management-go/app/configuration"
	"k8s-management-go/app/constants"
	"k8s-management-go/app/utils/cmdexecutor"
	"k8s-management-go/app/utils/files"
	"k8s-management-go/app/utils/logger"
	"k8s-management-go/app/utils/loggingstate"
)

// ExecuteScriptsInstallScriptsForNamespace executes install scripts, which have the prefix for the given namespace
func ExecuteScriptsInstallScriptsForNamespace(namespace string, filePrefix string) (err error) {
	log := logger.Log()
	log.Infof("[Execute Scripts] Try to execute scripts for namespace [%s]...", namespace)
	// calculate path to script folder
	var scriptFolder = files.AppendPath(
		files.AppendPath(
			configuration.GetConfiguration().GetProjectBaseDirectory(),
			namespace,
		),
		constants.DirProjectScripts,
	)

	// check if folder exists
	var isScriptsDirectoryAvailable = files.FileOrDirectoryExists(scriptFolder)
	if isScriptsDirectoryAvailable {
		log.Infof("[Execute Scripts] Script directory is available for namespace [%s]...", namespace)
		// prepare file filter for install
		scriptFileEnding := constants.ScriptsFileEnding
		var fileFilter = files.FileFilter{
			Prefix: &filePrefix,
			Suffix: &scriptFileEnding,
		}
		// list files which match to filter
		fileArray, err := files.ListFilesOfDirectoryWithFilter(scriptFolder, &fileFilter)
		if err != nil {
			log.Errorf("[Execute Scripts] Error while filtering for files with prefix [%s] in directory [%s].", filePrefix, scriptFolder)
			return err
		}

		// if array has content -> execute scripts
		if len(*fileArray) > 0 {
			// iterate over filtered file array and execute scripts
			for _, file := range *fileArray {
				scriptWithPath := files.AppendPath(scriptFolder, file)
				log.Infof("[Execute Scripts] Trying to execute script [%s]", scriptWithPath)
				loggingstate.AddInfoEntryAndDetails(fmt.Sprintf("  -> Try to execute script [%s]...", file), scriptWithPath)

				// Execute scripts
				_, _ = cmdexecutor.Executor.CombinedOutput("chmod", "755", scriptWithPath)
				outputCmd, err := cmdexecutor.Executor.CombinedOutput(scriptWithPath)
				if err != nil {
					loggingstate.AddErrorEntryAndDetails(fmt.Sprintf("  -> Try to execute script [%s]...failed. See output.", file), string(outputCmd))
					loggingstate.AddErrorEntryAndDetails(fmt.Sprintf("  -> Try to execute script [%s]...failed. See error.", file), err.Error())
					log.Errorf("Unable to execute script [%s] Error: \n%s", scriptWithPath, err.Error())
					log.Errorf("Unable to execute script [%s] Output: \n%s", scriptWithPath, string(outputCmd))

					return err
				}
				loggingstate.AddInfoEntryAndDetails(fmt.Sprintf("  -> Try to execute script [%s]...done. See output.", file), string(outputCmd))
				log.Infof("[Execute Scripts] Script output of [%s]: \n%s", scriptWithPath, outputCmd)
			}
		} else {
			loggingstate.AddInfoEntry(fmt.Sprintf("  -> No scripts with prefix [%s] found for [%s]", filePrefix, namespace))
		}
	} else {
		loggingstate.AddInfoEntry(fmt.Sprintf("  -> No scripts directory found for [%s]", namespace))
	}

	log.Infof("[Execute Scripts] Executing scripts for namespace [%s] done.", namespace)
	return nil
}
