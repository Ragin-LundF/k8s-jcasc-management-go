package validator

import "errors"

func ValidatePersistentVolumeClaim(input string) error {
	// a pvc name cannot be longer than 253 characters
	if len(input) > 253 {
		return errors.New("PVC name is too long! You can only use max. 253 characters. ")
	}
	return nil
}
