package project

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func TestCreateJenkinsHelmValues(t *testing.T) {
	testDefaultProjectConfiguration(t, false)

	var jenkinsHelmValues = NewJenkinsHelmValues()

	assertDefaultJenkinsHelmValues(jenkinsHelmValues, t)
}

func assertDefaultJenkinsHelmValues(jenkinsHelmValues *jenkinsHelmValues, t *testing.T) {
	assert.Equal(t, testJenkinsHelmMasterImage, jenkinsHelmValues.Master.Image)
	assert.Equal(t, testJenkinsHelmMasterImageTag, jenkinsHelmValues.Master.Tag)
	assert.Equal(t, testJenkinsHelmMasterPullPolicy, jenkinsHelmValues.Master.ImagePullPolicy)
	assert.Equal(t, testJenkinsHelmMasterPullSecret, jenkinsHelmValues.Master.ImagePullSecretName)
	assert.Equal(t, testJenkinsHelmMasterDefaultLabel, jenkinsHelmValues.Master.CustomJenkinsLabels)
	assert.Equal(t, testJenkinsHelmMasterAdminPassword, jenkinsHelmValues.Master.AdminPassword)
	assert.Equal(t, testJenkinsHelmMasterJcascConfigUrl, jenkinsHelmValues.Master.SidecarsConfigAutoReloadFolder)
	assert.Equal(t, strconv.FormatBool(!testJenkinsHelmMasterDenyAnonymousReadAccess), jenkinsHelmValues.Master.AuthorizationStrategyDenyAnonymousReadAccess)
	assert.Equal(t, testConfigJenkinsMasterPvcStorageClassName, jenkinsHelmValues.Persistence.StorageClass)
	assert.Equal(t, testConfigJenkinsMasterPvcAccessMode, jenkinsHelmValues.Persistence.AccessMode)
	assert.Equal(t, testConfigJenkinsMasterPvcSize, jenkinsHelmValues.Persistence.Size)
}
