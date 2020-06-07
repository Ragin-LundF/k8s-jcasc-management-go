package kubectl

import (
	"k8s-management-go/app/cli/loggingstate"
	"k8s-management-go/app/utils/logger"
	"os/exec"
	"strings"
)

// Execute Kubectl commands
func ExecutorKubectl(command string, args []string) (output string, err error) {
	log := logger.Log()

	// create args
	argsForCommand := []string{
		command,
	}

	// append args from method
	argsForCommand = append(argsForCommand, args...)

	loggingstate.AddInfoEntryAndDetails("  -> Executing K8S command...", "kubectl "+strings.Join(argsForCommand, " "))
	log.Info("[ExecKubectl] Executing K8S command: \n   -> kubectl %v", strings.Join(argsForCommand, " "))

	// execute
	cmdOutput, err := exec.Command("kubectl", argsForCommand...).CombinedOutput()
	if err != nil {
		// log output error
		loggingstate.AddErrorEntryAndDetails("  -> Unable to execute kubectl command. See output.", string(cmdOutput))
		loggingstate.AddErrorEntryAndDetails("  -> Unable to execute kubectl command. See error.", err.Error())
		log.Error("[ExecKubectl] -> K8S command failed. Output: \n%v", cmdOutput)
		log.Error("[ExecKubectl] -> K8S command failed. Error: \n%v", err)

		return string(cmdOutput), err
	}

	loggingstate.AddInfoEntryAndDetails("  -> Executing K8S command...done", string(cmdOutput))
	log.Info("[ExecKubectl] Executing K8S command done: \n%v", string(cmdOutput))

	return string(cmdOutput), err
}
