package encryption

import (
	"golang.org/x/crypto/bcrypt"
	"k8s-management-go/app/constants"
	"strings"
	"testing"
)

func TestEncryptJenkinsUserPassword(t *testing.T) {
	// prepare variables
	var password = "mypass"

	// executing the method
	result, err := EncryptJenkinsUserPassword(password)

	// validating the result
	if err != nil && result == "" {
		t.Errorf("Encryption exists with error: %s", err.Error())
	} else {
		encryptedPass := strings.TrimPrefix(result, constants.UtilsJenkinsUserPassBcryptPrefix)
		err = bcrypt.CompareHashAndPassword([]byte(encryptedPass), []byte(password))
		if err != nil {
			t.Error("Password is not comparable. Hash is maybe wrong.")
		} else {
			t.Log("Success validating password hash")
		}
	}
}
