package server

import (
	"encoding/json"
	"k8s-management-go/app/models"
	"net/http"
)

// ConfigurationAPI implements the API endpoint for the configuration
func ConfigurationAPI(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case "GET":
		configuration := models.GetConfiguration()
		configurationAsJSON, _ := json.MarshalIndent(configuration, "", "\t")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(configurationAsJSON))
	default:
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message": "not found"}`))
	}
}
