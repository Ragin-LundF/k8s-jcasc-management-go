package kubectl

import (
	"fmt"
	"k8s-management-go/app/utils/cmdexecutor"
	"k8s-management-go/app/utils/logger"
	"k8s-management-go/app/utils/loggingstate"
	"strings"
)

// ExecutorKubectl executes kubectl commands
func ExecutorKubectl(command string, args []string) (output string, err error) {
	var log = logger.Log()

	// create args
	var argsForCommand = []string{
		command,
	}

	// append args from method
	argsForCommand = append(argsForCommand, args...)

	loggingstate.AddInfoEntryAndDetails("  -> Executing K8S command...", fmt.Sprintf("kubectl %s", strings.Join(argsForCommand, " ")))
	log.Infof("[ExecKubectl] Executing K8S command: \n   -> kubectl %s", strings.Join(argsForCommand, " "))

	// execute
	cmdOutput, err := cmdexecutor.Executor.CombinedOutput("kubectl", argsForCommand...)
	if err != nil {
		// log output error
		loggingstate.AddErrorEntryAndDetails("  -> Unable to execute kubectl command. See output.", string(cmdOutput))
		loggingstate.AddErrorEntryAndDetails("  -> Unable to execute kubectl command. See error.", err.Error())
		log.Errorf("[ExecKubectl] -> K8S command failed. Output: \n%s", cmdOutput)
		log.Errorf("[ExecKubectl] -> K8S command failed. Error: \n%s", err.Error())

		return string(cmdOutput), err
	}

	loggingstate.AddInfoEntryAndDetails("  -> Executing K8S command...done", string(cmdOutput))
	log.Infof("[ExecKubectl] Executing K8S command done: \n%s", string(cmdOutput))

	return string(cmdOutput), err
}
