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
	log.Infof("[ExecHelm] Executing Helm command: \n%s", strings.Join(argsForCommand, " "))

	// execute
	cmdOutput, err := exec.Command("helm", argsForCommand...).CombinedOutput()
	if err != nil {
		// log output error
		loggingstate.AddErrorEntryAndDetails("[ExecHelm] -> Helm command failed. See details.", string(cmdOutput))
		loggingstate.AddErrorEntryAndDetails("[ExecHelm] -> Helm command failed. See errors.", err.Error())
		log.Errorf("[ExecHelm] -> Helm command failed. Output: \n%s", string(cmdOutput))
		log.Errorf("[ExecHelm] -> Helm command failed. Error: \n%s", err.Error())
		return err
	}
	loggingstate.AddInfoEntryAndDetails("   -> [ExecHelm] Executing Helm command...done", string(cmdOutput))
	log.Infof("   -> [ExecHelm] Executing Helm command...done. Output: \n%s", string(cmdOutput))

	return err
}
