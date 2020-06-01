package main

import (
	"k8s-management-go/cli"
	"k8s-management-go/utils"
)

func main() {
	// setup the system and read config
	utils.Setup()
	// start UI workflow
	cli.Workflow()
}
