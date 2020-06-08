package main

import (
	"k8s-management-go/app/cli"
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
	// start UI workflow
	cli.Workflow(info, nil)
}
