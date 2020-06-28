package validator

import "errors"

func JenkinsSystemMessageValidator(input string) error {
	// a namespace name cannot be longer than 63 characters
	if len(input) > 255 {
		return errors.New("Should not be longer than 255 characters. ")
	}
	return nil
}
