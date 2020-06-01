package api

import (
	"encoding/json"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

type JenkinsUserPassword struct {
	Password string `json:"password"`
}

func JenkinsUserPasswordApi(w http.ResponseWriter, r *http.Request) {
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
		hash, err := bcrypt.GenerateFromPassword([]byte(jenkinsUserPassword.Password), bcrypt.MinCost)
		if err != nil {
			log.Println(err)
			http.Error(w, `{"message": "cannot create bcrypted password"}`, http.StatusBadRequest)
			return
		}

		// assign bcrypt to new JenkinsUserPassword struct
		encryptedJenkinsUserPassword := JenkinsUserPassword{
			Password: "#jbcrypt:" + string(hash),
		}

		// marshall object to JSON
		encryptedJenkinsUserPasswordAsJson, err := json.Marshal(encryptedJenkinsUserPassword)
		if err != nil {
			log.Println(err)
			http.Error(w, `{"message": "error marhsalling the response"}`, http.StatusBadRequest)
			return
		}

		// write response if everything was ok
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(encryptedJenkinsUserPasswordAsJson)
	default:
		w.WriteHeader(http.StatusNotFound)
		_, _ = w.Write([]byte(`{"message": "not found"}`))
	}
}
