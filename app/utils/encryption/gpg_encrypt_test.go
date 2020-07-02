package encryption

import (
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

	if maskedString == strings.Join(gpgCmdArgsMasked, " ") {
		t.Log("Success. Password is masked.")
	} else {
		t.Errorf("Failed. Return string was: [%s], but should be [%s]", maskedString, strings.Join(gpgCmdArgsMasked, " "))
	}
}
