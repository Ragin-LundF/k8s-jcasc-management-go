package project

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func TestCreateJenkinsHelmValues(t *testing.T) {
	testDefaultProjectConfiguration(t, false)

	var jenkinsHelmValues = newJenkinsHelmValues()

	assertDefaultJenkinsHelmValues(jenkinsHelmValues, t)
}

func assertDefaultJenkinsHelmValues(jenkinsHelmValues *jenkinsHelmValues, t *testing.T) {
	assert.Equal(t, testJenkinsHelmMasterImage, jenkinsHelmValues.Controller.Image)
	assert.Equal(t, testJenkinsHelmMasterImageTag, jenkinsHelmValues.Controller.Tag)
	assert.Equal(t, testJenkinsHelmMasterPullPolicy, jenkinsHelmValues.Controller.ImagePullPolicy)
	assert.Equal(t, testJenkinsHelmMasterPullSecret, jenkinsHelmValues.Controller.ImagePullSecretName)
	assert.Equal(t, testJenkinsHelmMasterDefaultLabel, jenkinsHelmValues.Controller.CustomJenkinsLabels)
	assert.Equal(t, testJenkinsHelmMasterAdminPassword, jenkinsHelmValues.Controller.AdminPassword)
	assert.Equal(t, testJenkinsHelmMasterJcascConfigUrl, jenkinsHelmValues.Controller.SidecarsConfigAutoReloadFolder)
	assert.Equal(t, strconv.FormatBool(testJenkinsHelmMasterDenyAnonymousReadAccess), jenkinsHelmValues.Controller.AuthorizationStrategyAllowAnonymousRead)
	assert.Equal(t, testConfigJenkinsMasterPvcStorageClassName, jenkinsHelmValues.Persistence.StorageClass)
	assert.Equal(t, testConfigJenkinsMasterPvcAccessMode, jenkinsHelmValues.Persistence.AccessMode)
	assert.Equal(t, testConfigJenkinsMasterPvcSize, jenkinsHelmValues.Persistence.Size)
	assert.Equal(t, []string{}, jenkinsHelmValues.AdditionalNamespaces)
}
