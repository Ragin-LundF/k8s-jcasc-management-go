package server

import (
	"encoding/json"
	"k8s-management-go/app/models"
	"net/http"
)

// IPConfigurationAPI implements the API endpoint for getting the IP configuration
func IPConfigurationAPI(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case "GET":
		ipConfiguration := models.GetIPConfiguration()
		ipConfigurationAsJSON, _ := json.MarshalIndent(ipConfiguration, "", "\t")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(ipConfigurationAsJSON))
	default:
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message": "not found"}`))
	}
}
