package main

import (
	"k8s-management-go/api"
	"k8s-management-go/utils"
)

// Experimental server
func main() {
	utils.Setup()
	api.StartServer()
}
