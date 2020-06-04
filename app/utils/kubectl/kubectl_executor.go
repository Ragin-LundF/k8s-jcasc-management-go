package kubectl

import (
	"errors"
	"k8s-management-go/app/constants"
	"k8s-management-go/app/utils/logger"
	"os/exec"
	"strings"
)

// Execute Kubectl commands
func ExecutorKubectl(command string, args []string) (output string, info string, err error) {
	log := logger.Log()

	// create args
	argsForCommand := []string{
		command,
	}

	// append args from method
	argsForCommand = append(argsForCommand, args...)

	info = info + constants.NewLine + "[ExecKubectl] Executing K8S command:"
	info = info + constants.NewLine + "[ExecKubectl] kubectl " + strings.Join(argsForCommand, " ")
	log.Info(info)

	// execute
	cmdOutput, err := exec.Command("kubectl", argsForCommand...).CombinedOutput()
	if err != nil {
		// log output error
		log.Error("[ExecKubectl] -> K8S command failed. Output:")
		log.Error(string(cmdOutput) + constants.NewLine + err.Error())

		err = errors.New(string(cmdOutput) + constants.NewLine + err.Error())
	}

	return string(cmdOutput), info, err
}
