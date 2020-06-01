package main

import (
	"k8s-management-go/api"
	"k8s-management-go/models/config"
	"k8s-management-go/utils"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	setup()
	http.HandleFunc("/v1/k8smgmt/menu", api.MenuApi)
	http.HandleFunc("/v1/k8smgmt/configuration", api.ConfigurationApi)
	http.HandleFunc("/v1/k8smgmt/configuration/ip", api.IpConfigurationApi)
	http.HandleFunc("/v1/k8smgmt/jenkins/password", api.JenkinsUserPasswordApi)
	log.Fatal(http.ListenAndServe(":80", nil))
}

// initial setup
// reads the configuration
func setup() {
	// define main path
	basePath := ""
	if os.Getenv("K8S_MGMT_BASE_PATH") != "" {
		// base path from environment variables
		basePath = os.Getenv("K8S_MGMT_BASE_PATH")
	} else if len(os.Args) == 2 {
		// base path as argument found
		basePath = os.Args[1]
	}

	// read configuration if base path was set, else go into panic mode
	if basePath != "" {
		// prepare basePath
		basePath = strings.Replace(basePath, "\"", "", -1)
		basePath = strings.TrimSuffix(basePath, "/")

		// read configuration
		utils.ReadConfiguration(basePath)
		// receive configuration
		configuration := *config.GetConfiguration()
		// log configuration
		utils.LogStruct("Configuration", &configuration)
		// TODO: only for testing here
		utils.ReadIpConfig(basePath)
		utils.LogStruct("Ip Config", *config.GetIpConfiguration())
	} else {
		panic("No configuration path found...")
	}
}
