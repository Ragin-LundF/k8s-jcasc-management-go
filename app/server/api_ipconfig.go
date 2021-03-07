package server

import (
	"encoding/json"
	"k8s-management-go/app/configuration"
	"net/http"
)

// IPConfigurationAPI implements the API endpoint for getting the IP configuration
func IPConfigurationAPI(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case "GET":
		ipConfiguration := configuration.GetConfiguration().K8SManagement.IPConfig.Deployments
		ipConfigurationAsJSON, _ := json.MarshalIndent(ipConfiguration, "", "\t")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(ipConfigurationAsJSON))
	default:
		w.WriteHeader(http.StatusNotFound)
		_, _ = w.Write([]byte(`{"message": "not found"}`))
	}
}
