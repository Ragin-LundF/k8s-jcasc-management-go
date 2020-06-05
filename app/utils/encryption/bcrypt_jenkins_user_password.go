package encryption

import (
	"golang.org/x/crypto/bcrypt"
	"k8s-management-go/app/utils/logger"
)

func EncryptJenkinsUserPassword(plainPassword string) (hashedPassword string, err error) {
	log := logger.Log()

	// create bcrypt hash from password
	hashByte, err := bcrypt.GenerateFromPassword([]byte(plainPassword), bcrypt.MinCost)
	if err != nil {
		log.Error("[EncryptJenkinsUserPassword] Unable to encrypt password... %v\n", err)
		return "", err
	}
	return "#jbcrypt:" + string(hashByte), err
}
