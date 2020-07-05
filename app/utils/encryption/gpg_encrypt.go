package encryption

import (
	"errors"
	"fmt"
	"k8s-management-go/app/utils/cmdexecutor"
	"k8s-management-go/app/utils/files"
	"k8s-management-go/app/utils/loggingstate"
	"strings"
)

// GpgEncryptSecrets encrypts secrets
func GpgEncryptSecrets(secretsFilePath string, password string) (err error) {
	// first encrypt file
	gpgCmdArgs := []string{
		"--batch", "--yes", "--passphrase", password, "-c", secretsFilePath,
	}
	loggingstate.AddInfoEntryAndDetails("  -> Executing Gpg encrypt command...", fmt.Sprintf("gpg %s", maskedGpgCmdArgsAsString(gpgCmdArgs, 4)))

	cmdOutput, err := cmdexecutor.Executor.CombinedOutput("gpg", gpgCmdArgs...)
	if err != nil {
		loggingstate.AddErrorEntryAndDetails("  -> Unable to encrypt secrets file...failed. See output", string(cmdOutput))
		loggingstate.AddErrorEntryAndDetails("  -> Unable to encrypt secrets file...failed. See error", err.Error())
		return err
	}

	loggingstate.AddInfoEntry(fmt.Sprintf("  -> Encrypt secrets file [%s] done.", secretsFilePath))

	return nil
}

// GpgDecryptSecrets decrypts secrets
func GpgDecryptSecrets(secretsFilePath string, password string) (err error) {
	gpgCmdArgs := []string{
		"--batch", "--yes", "--passphrase", password, secretsFilePath,
	}
	loggingstate.AddInfoEntryAndDetails("  -> Executing Gpg decrypt command...", fmt.Sprintf("gpg %s", maskedGpgCmdArgsAsString(gpgCmdArgs, 4)))

	cmdOutput, err := cmdexecutor.Executor.CombinedOutput("gpg", gpgCmdArgs...)
	if err != nil {
		loggingstate.AddErrorEntryAndDetails("  -> Unable to decrypt secrets file...failed. See output", string(cmdOutput))
		loggingstate.AddErrorEntryAndDetails("  -> Unable to decrypt secrets file...failed. See error", err.Error())

		return err
	}

	// check if result is ok
	newSecretsFile := strings.Replace(secretsFilePath, ".gpg", "", -1)
	if !files.FileOrDirectoryExists(newSecretsFile) {
		loggingstate.AddErrorEntry(fmt.Sprintf("  -> Decrypting was not successful. Script file [%s] does not exist.", newSecretsFile))
		return errors.New("No decrypted secrets file found! ")
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
