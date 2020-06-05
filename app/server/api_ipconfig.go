package server

import (
	"encoding/json"
	"k8s-management-go/app/models"
	"net/http"
)

func IpConfigurationApi(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case "GET":
		ipConfiguration := models.GetIpConfiguration()
		ipConfigurationAsJson, _ := json.MarshalIndent(ipConfiguration, "", "\t")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(ipConfigurationAsJson))
	default:
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message": "not found"}`))
	}
}
