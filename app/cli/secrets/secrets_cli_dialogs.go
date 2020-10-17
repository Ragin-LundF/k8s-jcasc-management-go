package secrets

import (
	"errors"
	"k8s-management-go/app/actions/secretsactions"
	"k8s-management-go/app/cli/dialogs"
	"k8s-management-go/app/models"
	"k8s-management-go/app/utils/loggingstate"
	"strings"
)

// AskForSecretsPassword asks for password
func AskForSecretsPassword(passwordText string, selectSecretsFile bool) (secretsFile string, password string, err error) {
	if selectSecretsFile {
		// ask for secrets file
		loggingstate.AddInfoEntry("  -> Ask for secrets file to apply...")
		secretsFile, err = dialogs.DialogAskForSecretsFile()
		if err != nil {
			loggingstate.AddErrorEntryAndDetails("  -> Ask for secrets file to apply...failed", err.Error())
			return "", "", err
		}
		loggingstate.AddInfoEntry("  -> Ask for secrets file to apply...done")
	}

	// Validator for password (keep it simple for now)
	validate := func(input string) error {
		if len(input) < 4 {
			return errors.New("Password too short! ")
		}
		if strings.Contains(input, " ") {
			return errors.New("Password should not contain spaces! ")
		}
		return nil
	}

	// ask for password
	loggingstate.AddInfoEntry("  -> Ask for the password for secret file...")
	password, err = dialogs.DialogAskForPassword(passwordText, validate)
	if err != nil {
		loggingstate.AddErrorEntryAndDetails("  -> Ask for the password for secret files...failed", err.Error())

		return secretsFile, "", err
	}
	loggingstate.AddInfoEntry("  -> Ask for the password for secret file...done")

	return secretsFile, password, err
}

// ApplySecrets applies the secrets
func ApplySecrets() (err error) {
	// ask for namespace
	loggingstate.AddInfoEntry("  -> Ask for namespace to apply secrets...")

	namespace, err := dialogs.DialogAskForNamespace()
	if err != nil {
		loggingstate.AddErrorEntryAndDetails("  -> Ask for namespace to apply secrets...failed", err.Error())
		return err
	}
	loggingstate.AddInfoEntry("  -> Ask for namespace to apply secrets...done")

	// ask for secrets file
	loggingstate.AddInfoEntry("  -> Ask for secrets file to apply...")
	secretsFile, err := dialogs.DialogAskForSecretsFile()
	if err != nil {
		loggingstate.AddErrorEntryAndDetails("  -> Ask for secrets file to apply...failed", err.Error())
		return err
	}
	loggingstate.AddInfoEntry("  -> Ask for secrets file to apply...done")

	// apply secrets to namespace
	loggingstate.AddInfoEntry("  -> Ask for namespace to apply secrets...")

	if err = ApplySecretsToNamespace(namespace, secretsFile, nil); err != nil {
		loggingstate.AddErrorEntryAndDetails("  -> Ask for namespace to apply secrets...failed", err.Error())
		return err
	}

	loggingstate.AddInfoEntry("  -> Ask for namespace to apply secrets...done")

	return nil
}

// ApplySecretsToNamespace applies secrets to one namespace
func ApplySecretsToNamespace(namespace string, secretsFileName string, password *string) (err error) {
	// Decrypt secrets file
	if password != nil {
		if err = secretsactions.ActionDecryptSecretsFile(*password, secretsFileName); err != nil {
			return err
		}
	} else {
		if err = DecryptSecretsFile(&secretsFileName); err != nil {
			return err
		}
	}

	// apply secret to namespace
	err = secretsactions.ActionApplySecretsToNamespace(namespace, secretsFileName)
	return err
}

// ApplySecretsToAllNamespaces applies secrets to all namespaces
func ApplySecretsToAllNamespaces() (err error) {
	// ask for secrets file
	loggingstate.AddInfoEntry("  -> Ask for secrets file to apply...")
	secretsFile, err := dialogs.DialogAskForSecretsFile()
	if err != nil {
		loggingstate.AddErrorEntryAndDetails("  -> Ask for secrets file to apply...failed", err.Error())
		return err
	}
	loggingstate.AddInfoEntry("  -> Ask for secrets file to apply...done")

	err = DecryptSecretsFile(&secretsFile)
	if err != nil {
		return err
	}

	// prepare progressbar
	bar := dialogs.CreateProgressBar("Installing...", len(models.GetIPConfiguration().IPs))
	progress := dialogs.ProgressBar{
		Bar: &bar,
	}
	// apply secret to namespaces
	err = secretsactions.ActionApplySecretsToAllNamespaces(secretsFile, progress.AddCallback)

	return err
}
