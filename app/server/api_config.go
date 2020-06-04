package server

import (
	"encoding/json"
	"k8s-management-go/app/models/config"
	"net/http"
)

func ConfigurationApi(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case "GET":
		configuration := config.GetConfiguration()
		configurationAsJson, _ := json.MarshalIndent(configuration, "", "\t")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(configurationAsJson))
	default:
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message": "not found"}`))
	}
}
