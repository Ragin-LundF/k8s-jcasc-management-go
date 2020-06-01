package cli

import (
	"fmt"
	"github.com/manifoldco/promptui"
)

type confirm struct {
	Selection string
}

func ClearScreen() {
	fmt.Println("\033[2J")
}

func dialogConfirm(templateLabel string, templateSelector string, templateDetails string, dialogLabel string) bool {
	ClearScreen()

	// Template for displaying confirm dialog
	dialogConfim := []confirm{
		{Selection: "yes"},
		{Selection: "no"},
	}

	templates := &promptui.SelectTemplates{
		Label:    templateLabel,
		Active:   "\U000027A4 {{ ." + templateSelector + " | green }}",
		Inactive: "  {{ .Selection | cyan }}",
		Selected: "\U000027A4 {{ ." + templateSelector + " | red | cyan }}",
		Details:  templateDetails,
	}

	confirmDialogPrompt := promptui.Select{
		Label:     dialogLabel,
		Items:     dialogConfim,
		Templates: templates,
		Size:      2,
	}

	_, resultConfirm, err := confirmDialogPrompt.Run()

	if err != nil || resultConfirm != "{yes}" {
		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
		}
		return false
	} else {
		return true
	}
}

func dialogPassword(label string, validate promptui.ValidateFunc) (password string, err error) {
	ClearScreen()

	// Prepare prompt
	promptPlainPassword := promptui.Prompt{
		Label:    label,
		Validate: validate,
		Mask:     '*',
	}
	password, err = promptPlainPassword.Run()

	// check if everything was ok
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
	}
	return password, err
}
