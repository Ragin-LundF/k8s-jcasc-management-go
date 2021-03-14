package namespaceactions

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"k8s-management-go/app/actions/install"
	"k8s-management-go/app/utils/cmdexecutor"
	"strings"
	"testing"
)

func TestProcessNamespaceCreation(t *testing.T) {
	cmdexecutor.Executor = TestCommandExecKubectl{}
	var projectConfig = install.NewInstallProjectConfig()
	projectConfig.Project.SetNamespace("new-namespace")
	var err = ProcessNamespaceCreation(projectConfig)

	assert.NoError(t, err)
}

func TestProcessNamespaceCreationWithErr(t *testing.T) {
	cmdexecutor.Executor = TestCommandExecKubectlErr{}
	var projectConfig = install.NewInstallProjectConfig()
	projectConfig.Project.SetNamespace("new-namespace")
	var err = ProcessNamespaceCreation(projectConfig)

	assert.Error(t, err)
}

func TestIsNamespaceAvailable(t *testing.T) {
	cmdexecutor.Executor = TestCommandExecKubectl{}
	available, err := isNamespaceAvailable("project-b")

	assert.NoError(t, err)
	assert.True(t, available)
}

func TestIsNamespaceNotAvailable(t *testing.T) {
	cmdexecutor.Executor = TestCommandExecKubectl{}
	available, err := isNamespaceAvailable("project-not-existing")

	assert.NoError(t, err)
	assert.False(t, available)
}

func TestIsNamespaceWithError(t *testing.T) {
	cmdexecutor.Executor = TestCommandExecKubectlErr{}
	_, err := isNamespaceAvailable("project-not-existing")

	assert.Error(t, err)
}

// TestCommandExecKubectl is a mock with available namespace
type TestCommandExecKubectl struct{}

// TestCommandExecKubectlErr is a mock with error as result
type TestCommandExecKubectlErr struct{}

func (c TestCommandExecKubectl) CombinedOutput(command string, args ...string) ([]byte, error) {
	if command == "kubectl" {
		if strings.Join(args, " ") == "get namespaces" {
			return resultGetNamespaces()
		} else if len(args) == 3 && strings.Join(args, " ") == "create namespace "+args[2] {
			return resultCreateNamespace(args[2])
		}
	}
	return []byte("no result"), errors.New("No known command. ")
}

func (c TestCommandExecKubectlErr) CombinedOutput(command string, args ...string) ([]byte, error) {
	return nil, errors.New("exit with status 1")
}

// resultCreateNamespace generates output of expected kubectl create namespace xy
func resultCreateNamespace(namespace string) ([]byte, error) {
	return []byte("namespace/" + namespace + " created"), nil
}

// resultGetNamespaces generates output of expected kubectl get namespaces command
func resultGetNamespaces() ([]byte, error) {
	var kubectlResult = `NAME	STATUS	AGE
project-a	Active	10d
project-b	Active	10d
project-c	Active	10d
`
	return []byte(kubectlResult), nil
}
