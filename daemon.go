package main

import (
	"k8s-management-go/server"
	"k8s-management-go/utils"
)

// Experimental server
func main() {
	utils.Setup()
	server.StartServer()
}
