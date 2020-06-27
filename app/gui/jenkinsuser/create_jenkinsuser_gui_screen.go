package jenkinsuser

import (
	"fmt"
	"fyne.io/fyne"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
	"github.com/atotto/clipboard"
	"k8s-management-go/app/utils/encryption"
	"k8s-management-go/app/utils/loggingstate"
)

func ScreenJenkinsUserPasswordCreate(window fyne.Window) fyne.CanvasObject {
	// UI elements
	passwordErrorLabel := widget.NewLabel("")
	// secrets password
	userPasswordEntry := widget.NewPasswordEntry()
	userConfirmPasswordEntry := widget.NewPasswordEntry()

	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Password", Widget: userPasswordEntry},
			{Text: "Confirm password", Widget: userConfirmPasswordEntry},
			{Text: "", Widget: passwordErrorLabel},
		},
		OnSubmit: func() {
			isValid, errMessage := validatePasswords(userPasswordEntry.Text, userConfirmPasswordEntry.Text)
			passwordErrorLabel.SetText(errMessage)
			if isValid {
				hashedPassword, err := encryption.EncryptJenkinsUserPassword(userPasswordEntry.Text)
				if err != nil {
					dialog.NewError(err, window)
				}

				encPassEntry := widget.NewEntry()
				encPassEntry.Text = hashedPassword

				encPassBox := widget.NewVBox(
					widget.NewHBox(layout.NewSpacer()),
					encPassEntry,
					widget.NewHBox(layout.NewSpacer()),
					widget.NewLabelWithStyle("Do you want to copy the password to clipboard?", fyne.TextAlignLeading, fyne.TextStyle{Italic: true}),
				)
				dialog.NewCustomConfirm("Your encrypted password",
					"Copy it!",
					"No thanks!",
					encPassBox,
					func(wantCopy bool) {
						fmt.Println("Callback...")
						if wantCopy {
							fmt.Println("User wants to copy!")
							err = clipboard.WriteAll(hashedPassword)
							if err != nil {
								loggingstate.AddErrorEntryAndDetails("-> Unable to copy password to clipboard", err.Error())
							}
						}
						loggingstate.LogLoggingStateEntries()
					},
					window)
			}
		},
	}

	box := widget.NewVBox(
		widget.NewHBox(layout.NewSpacer()),
		form,
	)

	return box
}

func validatePasswords(password string, confirmPassword string) (isValid bool, errMessage string) {
	// check first, if both passwords are equal
	if password != confirmPassword {
		return false, "Passwords did not match!"
	}

	// check if password has a acceptable length (it is not enough, but better than nothing)
	if len(password) < 5 {
		return false, "Password length must be minimum 5 characters! Better will be more than 8!"
	}
	return true, ""
}
