package encryption

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

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
