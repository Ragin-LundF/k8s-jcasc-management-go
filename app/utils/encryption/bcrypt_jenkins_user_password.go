package encryption

import (
	"golang.org/x/crypto/bcrypt"
	"k8s-management-go/app/utils/loggingstate"
)

func EncryptJenkinsUserPassword(plainPassword string) (hashedPassword string, err error) {
	// create bcrypt hash from password
	hashByte, err := bcrypt.GenerateFromPassword([]byte(plainPassword), bcrypt.MinCost)
	if err != nil {
		loggingstate.AddErrorEntryAndDetails("  -> Unable to encrypt password.", err.Error())
		return "", err
	}
	return "#jbcrypt:" + string(hashByte), err
}
