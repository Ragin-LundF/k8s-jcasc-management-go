package server

import (
	"encoding/json"
	"k8s-management-go/app/configuration"
	"net/http"
)

// ConfigurationAPI implements the API endpoint for the configuration
func ConfigurationAPI(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case "GET":
		var config = configuration.GetConfiguration()
		configurationAsJSON, _ := json.MarshalIndent(config, "", "\t")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(configurationAsJSON))
	default:
		w.WriteHeader(http.StatusNotFound)
		_, _ = w.Write([]byte(`{"message": "not found"}`))
	}
}
