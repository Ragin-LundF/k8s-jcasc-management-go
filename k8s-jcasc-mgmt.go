package main

import (
	"k8s-management-go/app/cli"
	"k8s-management-go/app/utils/setup"
)

func main() {
	// setup the system and read config
	setup.Setup()
	// start UI workflow
	cli.Workflow("", nil)
}
