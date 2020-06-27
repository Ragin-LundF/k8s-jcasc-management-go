package encryption

import (
	"golang.org/x/crypto/bcrypt"
	"k8s-management-go/app/utils/logger"
	"k8s-management-go/app/utils/loggingstate"
)

func EncryptJenkinsUserPassword(plainPassword string) (hashedPassword string, err error) {
	log := logger.Log()

	// create bcrypt hash from password
	hashByte, err := bcrypt.GenerateFromPassword([]byte(plainPassword), bcrypt.MinCost)
	if err != nil {
		log.Errorf("[EncryptJenkinsUserPassword] Unable to encrypt password... %s\n", err.Error())
		loggingstate.AddErrorEntryAndDetails("  -> Unable to encrypt password.", err.Error())
		return "", err
	}
	return "#jbcrypt:" + string(hashByte), err
}
