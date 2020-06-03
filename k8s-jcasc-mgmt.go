package main

import (
	"k8s-management-go/app/cli"
	"k8s-management-go/app/utils"
)

func main() {
	// setup the system and read config
	utils.Setup()
	// start UI workflow
	cli.Workflow("", nil)
}
