package version

import (
	"github.com/stretchr/testify/assert"
	"regexp"
	"testing"
)

func TestCompareVersionsEqual(t *testing.T) {
	var localVersion = "1.4.0"
	var remoteVersion = "1.4.0"

	isLessThanRemote := compareVersions(localVersion, remoteVersion)

	assert.False(t, isLessThanRemote)
}

func TestCompareVersionsNewRemote(t *testing.T) {
	var localVersion = "1.4.0"
	var remoteVersion = "1.5.0"

	isLessThanRemote := compareVersions(localVersion, remoteVersion)

	assert.True(t, isLessThanRemote)
}

func TestCompareVersionsNewerLocal(t *testing.T) {
	var localVersion = "1.5.0"
	var remoteVersion = "1.4.0"

	isLessThanRemote := compareVersions(localVersion, remoteVersion)

	assert.False(t, isLessThanRemote)
}

func TestReceiveVersionFromGit(t *testing.T) {
	version, err := receiveVersionFromGit()

	assert.NoError(t, err)

	regex := regexp.MustCompile("^\\d*\\.\\d*.\\d*$")
	assert.True(t, regex.Match([]byte(version)))
}
