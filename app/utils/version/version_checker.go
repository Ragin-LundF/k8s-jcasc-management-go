package version

import (
	"bytes"
	"github.com/hashicorp/go-version"
	"golang.org/x/mod/semver"
	"io/ioutil"
	"k8s-management-go/app/utils/logger"
	"net/http"
)

// CheckVersion checks the version if there is a new one available
func CheckVersion() bool {
	var log = logger.Log()

	var remoteVersion, err = receiveVersionFromGit()
	if err != nil {
		log.Error(err)
		return false
	}
	localVersion, err := readLocalVersion()
	if err != nil {
		log.Error(err)
		return false
	}

	if remoteVersion == "" || localVersion == "" {
		return false
	}

	// check if version is valid to ensure that no 404 response throws errors later.
	// If the remote version is not valid, it returns true to avoid infinite no update info
	// if i.e. the repository was moved.
	if !semver.IsValid(remoteVersion) {
		return true
	}

	return compareVersions(localVersion, remoteVersion)
}

func compareVersions(localVersion string, remoteVersion string) bool {
	var semVerRemote, _ = version.NewSemver(remoteVersion)
	var semVerLocal, _ = version.NewSemver(localVersion)

	if semVerLocal.LessThan(semVerRemote) {
		return true
	}

	return false
}

func readLocalVersion() (version string, err error) {
	read, err := ioutil.ReadFile("VERSION")
	if err != nil {
		return "", err
	}
	return string(read), nil
}

func receiveVersionFromGit() (version string, err error) {
	resp, err := http.Get("https://raw.githubusercontent.com/Ragin-LundF/k8s-jcasc-management-go/main/VERSION")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var buffer = new(bytes.Buffer)
	_, err = buffer.ReadFrom(resp.Body)
	if err != nil {
		return "", err
	}
	version = string(buffer.Bytes())

	return version, nil
}
