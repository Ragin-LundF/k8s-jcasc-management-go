package encryption

import (
	"k8s-management-go/app/cli/loggingstate"
	"k8s-management-go/app/utils/logger"
	"os"
	"os/exec"
	"strings"
)

// encrypt secrets
func GpgEncryptSecrets(secretsFilePath string, password string) (err error) {
	log := logger.Log()
	// first encrypt file
	gpgCmdArgs := []string{
		"--batch", "--yes", "--passphrase", password, "-c", secretsFilePath,
	}
	loggingstate.AddInfoEntryAndDetails("  -> Executing Gpg encrypt command...", "gpg "+maskedGpgCmdArgsAsString(gpgCmdArgs, 4))
	log.Info("[GpgEncryptSecrets] Executing Gpg encrypt command: \n   -> gpg %v", maskedGpgCmdArgsAsString(gpgCmdArgs, 4))

	cmdOutput, err := exec.Command("gpg", gpgCmdArgs...).CombinedOutput()
	if err != nil {
		loggingstate.AddErrorEntryAndDetails("  -> Unable to encrypt secrets file...failed. See output", string(cmdOutput))
		loggingstate.AddErrorEntryAndDetails("  -> Unable to encrypt secrets file...failed. See error", err.Error())
		log.Error("[GpgEncryptSecrets] Unable to encrypt secrets file [%v]...Output: \n%v", secretsFilePath, string(cmdOutput))
		log.Error("[GpgEncryptSecrets] Unable to encrypt secrets file [%v]...Error: \n%v", secretsFilePath, err)
		return err
	}

	loggingstate.AddInfoEntry("  -> Encrypt secrets file [" + secretsFilePath + "] done.")
	log.Info("[GpgEncryptSecrets] Encrypt secrets file [%v] done.", secretsFilePath)

	// after everything was ok -> delete original file
	err = os.Remove(secretsFilePath)
	if err != nil {
		loggingstate.AddErrorEntryAndDetails("  -> Unable to delete decrypted secrets file ["+secretsFilePath+"].", err.Error())
		log.Error("[GpgEncryptSecrets] Unable to delete decrypted secrets file [%v].\n%v", secretsFilePath, err)
		return err
	}

	return nil
}

// decrypt secrets
func GpgDecryptSecrets(secretsFilePath string, password string) (err error) {
	log := logger.Log()
	gpgCmdArgs := []string{
		"--batch", "--yes", "--passphrase", password, secretsFilePath,
	}
	loggingstate.AddInfoEntryAndDetails("  -> Executing Gpg decrypt command...", "gpg "+maskedGpgCmdArgsAsString(gpgCmdArgs, 4))
	log.Info("[GpgDecryptSecrets] Executing Gpg decrypt command: \n   -> gpg %v", maskedGpgCmdArgsAsString(gpgCmdArgs, 4))

	cmdOutput, err := exec.Command("gpg", gpgCmdArgs...).CombinedOutput()
	if err != nil {
		loggingstate.AddErrorEntryAndDetails("  -> Unable to decrypt secrets file...failed. See output", string(cmdOutput))
		loggingstate.AddErrorEntryAndDetails("  -> Unable to decrypt secrets file...failed. See error", err.Error())
		log.Error("[GpgDecryptSecrets] Unable to decrypt secrets file [%v]...Output: \n%v", secretsFilePath, string(cmdOutput))
		log.Error("[GpgDecryptSecrets] Unable to decrypt secrets file [%v]...Error: \n%v", secretsFilePath, err)

		return err
	}

	loggingstate.AddInfoEntry("  -> Decrypt secrets file [" + secretsFilePath + "] done.")
	log.Info("[GpgDecryptSecrets] Decrypt secrets file [%v] done.", secretsFilePath)

	return nil
}

// output masked gpg cmd args
func maskedGpgCmdArgsAsString(gpgCmdArgs []string, passwordPos int8) string {
	maskedGpgCmdArgs := make([]string, len(gpgCmdArgs))
	copy(maskedGpgCmdArgs, gpgCmdArgs)
	maskedGpgCmdArgs[passwordPos-1] = "*****"
	return strings.Join(maskedGpgCmdArgs, " ")
}
