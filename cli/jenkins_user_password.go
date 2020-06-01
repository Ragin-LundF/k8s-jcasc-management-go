package cli

import (
	"errors"
	"fmt"
	"github.com/atotto/clipboard"
	"github.com/manifoldco/promptui"
	"k8s-management-go/utils/encryption"
	"log"
	"strings"
)

func CreateJenkinsUserPassword() (info string, err error) {
	// empty info
	info = ""
	// Validator for password (keep it simple for now)
	validate := func(input string) error {
		if len(input) < 4 {
			return errors.New("Password too short!")
		}
		if strings.Contains(input, " ") {
			return errors.New("Password should not contain spaces!")
		}
		return nil
	}

	// Prepare prompt
	promptPlainPassword := promptui.Prompt{
		Label:    "Plain Password",
		Validate: validate,
		Mask:     '*',
	}
	plainPassword, err := promptPlainPassword.Run()

	// check if everything was ok
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return info, err
	}

	// encrypt password with bcrypt
	hashedPassword, err := encryption.EncryptJenkinsUserPassword(plainPassword)
	if err != nil {
		log.Println(err)
		return info, err
	}

	templateDetails := `
--------- Encrypted Password ----------
{{ "Password    :" | faint }}	` + hashedPassword

	resultConfirm := dialogConfirm(
		"Do you want to copy the password to the clipboard?",
		"Selection",
		templateDetails,
		"Your password: "+hashedPassword,
	)

	if resultConfirm {
		// copy to clipboard
		err = clipboard.WriteAll(hashedPassword)
	}
	return "Created password: " + hashedPassword, err
}
