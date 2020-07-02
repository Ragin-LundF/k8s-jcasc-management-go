package version

import (
	"regexp"
	"testing"
)

func TestCompareVersionsEqual(t *testing.T) {
	var localVersion = "1.4.0"
	var remoteVersion = "1.4.0"

	isLessThanRemote := compareVersions(localVersion, remoteVersion)

	if isLessThanRemote {
		t.Error("Failed. Versions are equal, but compare did not recognize it.")
	} else {
		t.Log("Success. Both versions are equal.")
	}
}

func TestCompareVersionsNewRemote(t *testing.T) {
	var localVersion = "1.4.0"
	var remoteVersion = "1.5.0"

	isLessThanRemote := compareVersions(localVersion, remoteVersion)

	if isLessThanRemote {
		t.Log("Success. Compare recognizes greater remote.")
	} else {
		t.Error("Failed. Remote version is newer, but compare did not recognize it.")
	}
}

func TestCompareVersionsNewerLocal(t *testing.T) {
	var localVersion = "1.5.0"
	var remoteVersion = "1.4.0"

	isLessThanRemote := compareVersions(localVersion, remoteVersion)

	if isLessThanRemote {
		t.Error("Failed. Local version is newer, but compare did not recognize it.")
	} else {
		t.Log("Success. Compare recognized that local is newer than remote.")
	}
}

func TestReceiveVersionFromGit(t *testing.T) {
	version, err := receiveVersionFromGit()

	if err != nil {
		t.Error("Failed. Can not receive version from Git")
	}

	regex := regexp.MustCompile("^\\d*\\.\\d*.\\d*$")
	if regex.Match([]byte(version)) {
		t.Log("Success. Received a valid version")
	} else {
		t.Errorf("Failed. Received [%v], which is not a valid version.", version)
	}
}
