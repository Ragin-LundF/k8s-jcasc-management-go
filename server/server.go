package server

import (
	"log"
	"net/http"
)

// Experimental server
func StartServer() {
	// register API
	http.HandleFunc("/v1/k8smgmt/menu", MenuApi)
	http.HandleFunc("/v1/k8smgmt/configuration", ConfigurationApi)
	http.HandleFunc("/v1/k8smgmt/configuration/ip", IpConfigurationApi)
	http.HandleFunc("/v1/k8smgmt/jenkins/password", JenkinsUserPasswordApi)
	log.Fatal(http.ListenAndServe(":80", nil))
}
