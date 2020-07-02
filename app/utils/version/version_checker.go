package version

import (
	"bytes"
	"github.com/hashicorp/go-version"
	"io/ioutil"
	"k8s-management-go/app/utils/logger"
	"net/http"
)

func CheckVersion() bool {
	log := logger.Log()

	remoteVersion, err := receiveVersionFromGit()
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

	semVerRemote, _ := version.NewSemver(remoteVersion)
	semVerLocal, _ := version.NewSemver(localVersion)

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
	resp, err := http.Get("https://raw.githubusercontent.com/Ragin-LundF/k8s-jcasc-management-go/master/VERSION")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	buffer := new(bytes.Buffer)
	_, err = buffer.ReadFrom(resp.Body)
	if err != nil {
		return "", err
	}
	version = string(buffer.Bytes())

	return version, nil
}
