package encryption

import (
	"golang.org/x/crypto/bcrypt"
)

func EncryptJenkinsUserPassword(plainPassword string) (hashedPassword string, err error) {
	// create bcrypt hash from password
	hashByte, err := bcrypt.GenerateFromPassword([]byte(plainPassword), bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	return "#jbcrypt:" + string(hashByte), err
}
