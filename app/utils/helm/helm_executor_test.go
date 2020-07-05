package helm

import (
	"k8s-management-go/app/utils/cmdexecutor"
	"strings"
	"testing"
)

// TestCommandExec is the test executor for mocks
type TestCommandExec struct{}

func (c TestCommandExec) CombinedOutput(command string, args ...string) ([]byte, error) {
	var commandAsString = command + " " + strings.Join(args, " ")
	return []byte(commandAsString + "...executed"), nil
}

// tests the method without executing the helm command
func TestExecutorHelm(t *testing.T) {
	cmdexecutor.Executor = TestCommandExec{}
	arr := []string{"-dry-run"}
	err := ExecutorHelm("install", arr)
	if err != nil {
		t.Errorf("Failed. %s", err.Error())
	}
	t.Log("Fine")
}
