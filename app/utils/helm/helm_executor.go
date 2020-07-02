package helm

import (
	"fmt"
	"k8s-management-go/app/utils/loggingstate"
	"os/exec"
	"strings"
)

// ExecutorHelm executes Helm commands
func ExecutorHelm(command string, args []string) (err error) {
	// create args
	argsForCommand := []string{
		command,
	}

	// append args from method
	argsForCommand = append(argsForCommand, args...)

	loggingstate.AddInfoEntryAndDetails("   -> [ExecHelm] Executing Helm command...", fmt.Sprintf("helm %s", strings.Join(argsForCommand, " ")))

	// execute
	cmdOutput, err := exec.Command("helm", argsForCommand...).CombinedOutput()
	if err != nil {
		// log output error
		loggingstate.AddErrorEntryAndDetails("[ExecHelm] -> Helm command failed. See details.", string(cmdOutput))
		loggingstate.AddErrorEntryAndDetails("[ExecHelm] -> Helm command failed. See errors.", err.Error())
		return err
	}
	loggingstate.AddInfoEntryAndDetails("   -> [ExecHelm] Executing Helm command...done", string(cmdOutput))

	return err
}
