package secrets

import (
	"k8s-management-go/app/actions/secrets_actions"
	"k8s-management-go/app/cli/dialogs"
	"k8s-management-go/app/models"
	"k8s-management-go/app/utils/loggingstate"
)

func ApplySecrets() (err error) {
	// ask for namespace
	loggingstate.AddInfoEntry("  -> Ask for namespace to apply secrets...")

	namespace, err := dialogs.DialogAskForNamespace()
	if err != nil {
		loggingstate.AddErrorEntryAndDetails("  -> Ask for namespace to apply secrets...failed", err.Error())
		return err
	}

	loggingstate.AddInfoEntry("  -> Ask for namespace to apply secrets...done")

	// apply secrets to namespace
	loggingstate.AddInfoEntry("  -> Ask for namespace to apply secrets...")

	if err = ApplySecretsToNamespace(namespace, nil); err != nil {
		loggingstate.AddErrorEntryAndDetails("  -> Ask for namespace to apply secrets...failed", err.Error())
		return err
	}

	loggingstate.AddInfoEntry("  -> Ask for namespace to apply secrets...done")

	return nil
}

// apply secrets to one namespace
func ApplySecretsToNamespace(namespace string, password *string) (err error) {
	// Decrypt secrets file
	if password != nil {
		if err = secrets_actions.ActionDecryptSecretsFile(*password); err != nil {
			return err
		}
	} else {
		if err = DecryptSecretsFile(); err != nil {
			return err
		}
	}

	// apply secret to namespace
	err = secrets_actions.ActionApplySecretsToNamespace(namespace)
	return err
}

// apply secrets to all namespaces
func ApplySecretsToAllNamespaces() (err error) {
	// apply secret to namespaces
	for _, ip := range models.GetIpConfiguration().Ips {
		if err := secrets_actions.ActionApplySecretsToNamespace(ip.Namespace); err != nil {
			return err
		}
	}
	return nil
}
