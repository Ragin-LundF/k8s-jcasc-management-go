package jenkinsuser

import (
	"github.com/atotto/clipboard"
	"k8s-management-go/app/cli/dialogs"
	"k8s-management-go/app/utils/encryption"
	"k8s-management-go/app/utils/loggingstate"
)

// CreateJenkinsUserPassword creates the Jenkins user password CLI dialog
func CreateJenkinsUserPassword() (err error) {
	loggingstate.AddInfoEntry("-> Ask for plain passwords...")

	plainPassword, err := ShowJenkinsUserPasswordDialog()
	if err != nil {
		return err
	}

	// compare plain passwords
	loggingstate.AddInfoEntry("  -> Entered passwords did match...Try to encrypt...")
	// encrypt password with bcrypt
	hashedPassword, err := encryption.EncryptJenkinsUserPassword(*plainPassword)
	if err != nil {
		return err
	}
	loggingstate.AddInfoEntry("-> Password successfully encrypted")

	templateDetails := `
--------- Encrypted Password ----------
{{ "Password    :" | faint }}	` + hashedPassword

	resultConfirm := dialogs.DialogConfirm(
		"Do you want to copy the password to the clipboard?",
		"Selection",
		templateDetails,
		"Your password: "+hashedPassword,
	)

	if resultConfirm {
		// copy to clipboard
		err = clipboard.WriteAll(hashedPassword)
		if err != nil {
			loggingstate.AddErrorEntryAndDetails("-> Unable to copy password to clipboard", err.Error())
		}
	}
	return err
}
