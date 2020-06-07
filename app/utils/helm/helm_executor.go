package helm

import (
	"k8s-management-go/app/cli/loggingstate"
	"k8s-management-go/app/utils/logger"
	"os/exec"
	"strings"
)

// Execute Helm commands
func ExecutorHelm(command string, args []string) (err error) {
	log := logger.Log()

	// create args
	argsForCommand := []string{
		command,
	}

	// append args from method
	argsForCommand = append(argsForCommand, args...)

	loggingstate.AddInfoEntryAndDetails("   -> [ExecHelm] Executing Helm command...", strings.Join(argsForCommand, " "))
	log.Info("[ExecHelm] Executing Helm command: \n%v", strings.Join(argsForCommand, " "))

	// execute
	cmdOutput, err := exec.Command("helm", argsForCommand...).CombinedOutput()
	if err != nil {
		// log output error
		loggingstate.AddErrorEntryAndDetails("[ExecHelm] -> Helm command failed. See details.", string(cmdOutput))
		loggingstate.AddErrorEntryAndDetails("[ExecHelm] -> Helm command failed. See errors.", err.Error())
		log.Error("[ExecHelm] -> Helm command failed. Output: \n%v", string(cmdOutput))
		log.Error("[ExecHelm] -> Helm command failed. Error: \n%v", err.Error())
		return err
	}
	loggingstate.AddInfoEntry("   -> [ExecHelm] Executing Helm command...done")

	return err
}
