package cli

import (
	"errors"
	"fmt"
	"github.com/atotto/clipboard"
	"github.com/manifoldco/promptui"
	"k8s-management-go/utils"
	"log"
	"strings"
)

type confirm struct {
	Selection string
}

func CreateJenkinsUserPassword() {
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
	}
	plainPassword, err := promptPlainPassword.Run()

	// check if everything was ok
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	// encrypt password with bcrypt
	hashedPassword, err := utils.EncryptJenkinsUserPassword(plainPassword)
	if err != nil {
		log.Println(err)
		return
	}

	// clear screen
	fmt.Println("\033[2J")

	// Template for displaying confirm dialog
	dialogConfim := []confirm{
		{Selection: "yes"},
		{Selection: "no"},
	}

	templates := &promptui.SelectTemplates{
		Label:    "Do you want to copy the password to the clipboard?",
		Active:   "\U000027A4 {{ .Selection | green }}",
		Inactive: "  {{ .Selection | cyan }}",
		Selected: "\U000027A4 {{ .Selection | red | cyan }}",
		Details: `
--------- Encrypted Password ----------
{{ "Password    :" | faint }}	` + hashedPassword,
	}

	confirmDialogPrompt := promptui.Select{
		Label:     "Your password: {{ hashedPassword }}",
		Items:     dialogConfim,
		Templates: templates,
		Size:      2,
	}

	_, resultConfirm, err := confirmDialogPrompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	if resultConfirm == "{yes}" {
		// copy to clipboard
		_ = clipboard.WriteAll(hashedPassword)
	}
}
