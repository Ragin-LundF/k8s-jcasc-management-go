package server

import (
	"log"
	"net/http"
)

// StartServer starts the experimental server
func StartServer() {
	// register API
	http.HandleFunc("/v1/k8smgmt/menu", MenuAPI)
	http.HandleFunc("/v1/k8smgmt/configuration", ConfigurationAPI)
	http.HandleFunc("/v1/k8smgmt/configuration/ip", IPConfigurationAPI)
	http.HandleFunc("/v1/k8smgmt/jenkins/password", JenkinsUserPasswordAPI)
	log.Fatal(http.ListenAndServe(":80", nil))
}
