package version

import (
	"github.com/stretchr/testify/assert"
	"regexp"
	"testing"
)

func TestCompareVersionsEqual(t *testing.T) {
	const localVersion = "1.4.0"
	const remoteVersion = "1.4.0"

	var isLessThanRemote = compareVersions(localVersion, remoteVersion)

	assert.False(t, isLessThanRemote)
}

func TestCompareVersionsNewRemote(t *testing.T) {
	const localVersion = "1.4.0"
	const remoteVersion = "1.5.0"

	var isLessThanRemote = compareVersions(localVersion, remoteVersion)

	assert.True(t, isLessThanRemote)
}

func TestCompareVersionsNewerLocal(t *testing.T) {
	const localVersion = "1.5.0"
	const remoteVersion = "1.4.0"

	var isLessThanRemote = compareVersions(localVersion, remoteVersion)

	assert.False(t, isLessThanRemote)
}

func TestReceiveVersionFromGit(t *testing.T) {
	var version, err = receiveVersionFromGit()

	assert.NoError(t, err)

	var regex = regexp.MustCompile("^\\d*\\.\\d*.\\d*$")
	assert.True(t, regex.Match([]byte(version)))
}
