package jenkinsuser

import (
	"errors"
	"k8s-management-go/app/cli/dialogs"
	"k8s-management-go/app/utils/loggingstate"
	"strings"
)

func ShowJenkinsUserPasswordDialog() (err error, password *string) {
	// Validator for password (keep it simple for now)
	validate := func(input string) error {
		if len(input) < 4 {
			return errors.New("Password too short! ")
		}
		if strings.Contains(input, " ") {
			return errors.New("Password should not contain spaces!")
		}
		return nil
	}

	// Get plain passwords
	plainPassword, err := dialogs.DialogAskForPassword("Password", validate)
	if err != nil {
		loggingstate.AddErrorEntryAndDetails("  -> Unable to get plain password.", err.Error())
		return err, nil
	}
	plainPasswordConfirm, err := dialogs.DialogAskForPassword("Retype your password", validate)
	if err != nil {
		loggingstate.AddErrorEntryAndDetails("  -> Unable to get plain confirmation password.", err.Error())
		return err, nil
	}

	if plainPassword != plainPasswordConfirm {
		loggingstate.AddErrorEntry("  -> Passwords did not match! Aborting...")
		return errors.New("Passwords did not match! "), nil
	}
	return err, &plainPassword
}
