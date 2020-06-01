package server

import (
	"k8s-management-go/models/config"
	"k8s-management-go/utils"
	"log"
	"net/http"
)

// Experimental server
func StartServer() {
	// receive configuration
	configuration := *config.GetConfiguration()
	// log configuration
	utils.LogStruct("Configuration", &configuration)

	// register API
	http.HandleFunc("/v1/k8smgmt/menu", MenuApi)
	http.HandleFunc("/v1/k8smgmt/configuration", ConfigurationApi)
	http.HandleFunc("/v1/k8smgmt/configuration/ip", IpConfigurationApi)
	http.HandleFunc("/v1/k8smgmt/jenkins/password", JenkinsUserPasswordApi)
	log.Fatal(http.ListenAndServe(":80", nil))
}
