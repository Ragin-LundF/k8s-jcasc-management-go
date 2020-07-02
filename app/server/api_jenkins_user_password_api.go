package server

import (
	"encoding/json"
	"k8s-management-go/app/utils/encryption"
	"log"
	"net/http"
)

// JenkinsUserPassword defines the JSON for the password
type JenkinsUserPassword struct {
	Password string `json:"password"`
}

// JenkinsUserPasswordAPI is an API for encrypting the Jenkins user password
func JenkinsUserPasswordAPI(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case "POST":
		var jenkinsUserPassword JenkinsUserPassword

		// Try to decode the request body into the struct. If there is an error,
		// respond to the client with the error message and a 400 status code.
		err := json.NewDecoder(r.Body).Decode(&jenkinsUserPassword)
		if err != nil || jenkinsUserPassword.Password == "" {
			log.Println(err)
			http.Error(w, `{"message": "invalid data"}`, http.StatusBadRequest)
			return
		}

		// create bcrypt hash from password
		hashedPassword, err := encryption.EncryptJenkinsUserPassword(jenkinsUserPassword.Password)
		if err != nil {
			log.Println(err)
			http.Error(w, `{"message": "cannot create bcrypted password"}`, http.StatusBadRequest)
			return
		}

		// assign bcrypt to new JenkinsUserPassword struct
		encryptedJenkinsUserPassword := JenkinsUserPassword{
			Password: hashedPassword,
		}

		// marshall object to JSON
		encryptedJenkinsUserPasswordAsJSON, err := json.Marshal(encryptedJenkinsUserPassword)
		if err != nil {
			log.Println(err)
			http.Error(w, `{"message": "error marhsalling the response"}`, http.StatusBadRequest)
			return
		}

		// write response if everything was ok
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(encryptedJenkinsUserPasswordAsJSON)
	default:
		w.WriteHeader(http.StatusNotFound)
		_, _ = w.Write([]byte(`{"message": "not found"}`))
	}
}
