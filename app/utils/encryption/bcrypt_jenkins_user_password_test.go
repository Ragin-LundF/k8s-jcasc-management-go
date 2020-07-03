package encryption

import (
	"github.com/stretchr/testify/assert"
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
	assert.NoError(t, err)
	assert.NotEqual(t, "", result)

	encryptedPass := strings.TrimPrefix(result, constants.UtilsJenkinsUserPassBcryptPrefix)
	err = bcrypt.CompareHashAndPassword([]byte(encryptedPass), []byte(password))
	assert.NoError(t, err)
}
