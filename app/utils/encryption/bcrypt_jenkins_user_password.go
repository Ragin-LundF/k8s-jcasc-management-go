package encryption

import (
	"golang.org/x/crypto/bcrypt"
	"k8s-management-go/app/constants"
	"k8s-management-go/app/utils/loggingstate"
)

// EncryptJenkinsUserPassword encrypts a plain password with bcrypt
func EncryptJenkinsUserPassword(plainPassword string) (hashedPassword string, err error) {
	// create bcrypt hash from password
	hashByte, err := bcrypt.GenerateFromPassword([]byte(plainPassword), bcrypt.MinCost)
	if err != nil {
		loggingstate.AddErrorEntryAndDetails("  -> Unable to encrypt password.", err.Error())
		return "", err
	}
	return constants.UtilsJenkinsUserPassBcryptPrefix + string(hashByte), err
}
