package cmdexecutor

import (
	"os/exec"
)

// CommandExec is an interface for executing exec.Command()
type CommandExec interface {
	// CombinedOutput is an interface copy of exec.Command().CombinedOutput()
	CombinedOutput(string, ...string) ([]byte, error)
}

// OsCommandExec is the real OS executor
type OsCommandExec struct{}

// Executor is the global var to set the CommandExec implementation
var Executor CommandExec

// CombinedOutput is executes exec.Command(command, args...).CombinedOutput()
func (c OsCommandExec) CombinedOutput(command string, args ...string) ([]byte, error) {
	return exec.Command(command, args...).CombinedOutput()
}
