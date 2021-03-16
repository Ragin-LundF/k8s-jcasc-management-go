package main

import (
	"k8s-management-go/app"
	"k8s-management-go/app/configuration"
	"k8s-management-go/app/utils/setup"
	"k8s-management-go/app/utils/version"
)

func main() {
	// setup the system and read config
	setup.Setup()
	// check version
	var info = checkVersion()

	// start app
	startApp(info)
}

func checkVersion() string {
	newVersionAvailable := version.CheckVersion()
	var info = ""
	if newVersionAvailable {
		info = "A new version is available!"
	}

	return info
}

func startApp(info string) {
	if configuration.GetConfiguration().K8SManagement.CliOnly {
		// cli
		app.StartCli(info)
	} else {
		// start UI workflow
		app.StartApp(info)
	}
}
