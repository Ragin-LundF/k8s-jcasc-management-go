package menu

import (
	"fyne.io/fyne/test"
	"github.com/stretchr/testify/assert"
	"k8s-management-go/app/utils/cmdexecutor"
	"strings"
	"testing"
)

func TestCreateTabMenu(t *testing.T) {
	cmdexecutor.Executor = TestCommandExecKubectl{}
	app := test.NewApp()
	window := app.NewWindow("test")
	tabMenu := CreateTabMenu(app, window, "")

	assert.Equal(t, "Welcome", tabMenu.Items[0].Text)
	assert.Equal(t, "Deployments", tabMenu.Items[1].Text)
	assert.Equal(t, "Secrets", tabMenu.Items[2].Text)
	assert.Equal(t, "Create Project", tabMenu.Items[3].Text)
	assert.Equal(t, "Tools", tabMenu.Items[4].Text)
}

// TestCommandExecKubectl is a mock with available namespace
type TestCommandExecKubectl struct{}

func (c TestCommandExecKubectl) CombinedOutput(command string, args ...string) ([]byte, error) {
	if command == "kubectl" {
		if strings.Join(args, " ") == "config current-context" {
			return []byte("my-ctx"), nil
		} else if strings.Join(args, " ") == "config get-context" {
			return resultGetContexts()
		}
	}
	return []byte("no result"), nil
}

// resultGetNamespaces generates output of expected kubectl get namespaces command
func resultGetContexts() ([]byte, error) {
	var kubectlResult = `CURRENT   NAME          CLUSTER   AUTHINFO   NAMESPACE
*         k8s-ds        k8s       default
          scratch-ctx   scratch   default
`
	return []byte(kubectlResult), nil
}
