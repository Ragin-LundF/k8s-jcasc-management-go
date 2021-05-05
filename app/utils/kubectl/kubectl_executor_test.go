package kubectl

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"k8s-management-go/app/utils/cmdexecutor"
	"testing"
)

func TestExecutorKubectl(t *testing.T) {
	cmdexecutor.Executor = TestCommandExecKubectl{}
	var args = []string{
		"get",
		"namespace",
	}
	var result, err = ExecutorKubectl("kubectl", args)
	assert.NotNil(t, result)
	assert.Nil(t, err)
	assert.Equal(t, "works", result)
}

func TestExecutorKubectlErr(t *testing.T) {
	cmdexecutor.Executor = TestCommandExecKubectlErr{}
	var args = []string{
		"got",
		"namespace",
	}
	var result, err = ExecutorKubectl("kubectl", args)
	assert.Empty(t, result)
	assert.NotNil(t, err)
	assert.Equal(t, "exit with status 1", err.Error())
}

// TestCommandExecKubectl is a mock with available namespace
type TestCommandExecKubectl struct{}

// TestCommandExecKubectlErr is a mock with error as result
type TestCommandExecKubectlErr struct{}

func (c TestCommandExecKubectl) CombinedOutput(command string, args ...string) ([]byte, error) {
	return []byte("works"), nil
}

func (c TestCommandExecKubectlErr) CombinedOutput(command string, args ...string) ([]byte, error) {
	return nil, errors.New("exit with status 1")
}
