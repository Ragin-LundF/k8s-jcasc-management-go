package encryption

import (
	"fmt"
	"k8s-management-go/app/utils/loggingstate"
	"os"
	"os/exec"
	"strings"
)

// encrypt secrets
func GpgEncryptSecrets(secretsFilePath string, password string) (err error) {
	// first encrypt file
	gpgCmdArgs := []string{
		"--batch", "--yes", "--passphrase", password, "-c", secretsFilePath,
	}
	loggingstate.AddInfoEntryAndDetails("  -> Executing Gpg encrypt command...", fmt.Sprintf("gpg %s", maskedGpgCmdArgsAsString(gpgCmdArgs, 4)))

	cmdOutput, err := exec.Command("gpg", gpgCmdArgs...).CombinedOutput()
	if err != nil {
		loggingstate.AddErrorEntryAndDetails("  -> Unable to encrypt secrets file...failed. See output", string(cmdOutput))
		loggingstate.AddErrorEntryAndDetails("  -> Unable to encrypt secrets file...failed. See error", err.Error())
		return err
	}

	loggingstate.AddInfoEntry(fmt.Sprintf("  -> Encrypt secrets file [%s] done.", secretsFilePath))

	// after everything was ok -> delete original file
	err = os.Remove(secretsFilePath)
	if err != nil {
		loggingstate.AddErrorEntryAndDetails(fmt.Sprintf("  -> Unable to delete decrypted secrets file [%s].", secretsFilePath), err.Error())
		return err
	}

	return nil
}

// decrypt secrets
func GpgDecryptSecrets(secretsFilePath string, password string) (err error) {
	gpgCmdArgs := []string{
		"--batch", "--yes", "--passphrase", password, secretsFilePath,
	}
	loggingstate.AddInfoEntryAndDetails("  -> Executing Gpg decrypt command...", fmt.Sprintf("gpg %s", maskedGpgCmdArgsAsString(gpgCmdArgs, 4)))

	cmdOutput, err := exec.Command("gpg", gpgCmdArgs...).CombinedOutput()
	if err != nil {
		loggingstate.AddErrorEntryAndDetails("  -> Unable to decrypt secrets file...failed. See output", string(cmdOutput))
		loggingstate.AddErrorEntryAndDetails("  -> Unable to decrypt secrets file...failed. See error", err.Error())

		return err
	}

	loggingstate.AddInfoEntry(fmt.Sprintf("  -> Decrypt secrets file [%s] done.", secretsFilePath))

	return nil
}

// output masked gpg cmd args
func maskedGpgCmdArgsAsString(gpgCmdArgs []string, passwordPos int8) string {
	maskedGpgCmdArgs := make([]string, len(gpgCmdArgs))
	copy(maskedGpgCmdArgs, gpgCmdArgs)
	maskedGpgCmdArgs[passwordPos-1] = "*****"
	return strings.Join(maskedGpgCmdArgs, " ")
}
