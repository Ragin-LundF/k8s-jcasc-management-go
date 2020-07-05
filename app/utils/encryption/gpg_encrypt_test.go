package encryption

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"k8s-management-go/app/utils/cmdexecutor"
	"strings"
	"testing"
)

// TestCommandExec is the test executor for mocks
type TestCommandExec struct{}

// TestCommandExecErr is the test executor for mocks with error
type TestCommandExecErr struct{}

func (c TestCommandExec) CombinedOutput(command string, args ...string) ([]byte, error) {
	var commandAsString = command + " " + strings.Join(args, " ")
	return []byte(commandAsString + "...executed"), nil
}

func (c TestCommandExecErr) CombinedOutput(command string, args ...string) ([]byte, error) {
	return []byte("Exit with status 1"), errors.New("Error ")
}

func TestGpgEncryptSecrets(t *testing.T) {
	cmdexecutor.Executor = TestCommandExec{}
	err := GpgEncryptSecrets("/tmp/secrets.sh", "password")

	assert.NoError(t, err)
}

func TestGpgEncryptSecretsWithErr(t *testing.T) {
	cmdexecutor.Executor = TestCommandExecErr{}
	err := GpgEncryptSecrets("/tmp/secrets.sh", "password")

	assert.Error(t, err)
}

func TestMaskedGpgCmdArgsAsString(t *testing.T) {
	// first encrypt file
	var gpgCmdArgs = []string{
		"--batch", "--yes", "--passphrase", "my_password_was_set_here", "-c", "/path/to/secrets/file",
	}
	var gpgCmdArgsMasked = []string{
		"--batch", "--yes", "--passphrase", "*****", "-c", "/path/to/secrets/file",
	}
	maskedString := maskedGpgCmdArgsAsString(gpgCmdArgs, 4)

	assert.Equal(t, maskedString, strings.Join(gpgCmdArgsMasked, " "))
}
