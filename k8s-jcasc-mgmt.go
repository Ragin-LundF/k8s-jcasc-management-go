package main

import (
	"k8s-management-go/app"
	"k8s-management-go/app/models"
	"k8s-management-go/app/utils/setup"
	"k8s-management-go/app/utils/version"
)

func main() {
	// setup the system and read config
	setup.Setup()
	// check version
	newVersionAvailable := version.CheckVersion()
	info := ""
	if newVersionAvailable {
		info = "A new version is available!"
	}

	if models.GetConfiguration().CliOnly {
		// cli
		app.StartCli(info)
	} else {
		// start UI workflow
		app.StartApp(info)
	}
}
