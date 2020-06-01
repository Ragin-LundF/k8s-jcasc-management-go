package api

import (
	"encoding/json"
	"net/http"
)

type MenuElement struct {
	Name        string
	Description string
}

type Menu struct {
	Elements []MenuElement
}

func MenuApi(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case "GET":
		menuAsJson, _ := json.MarshalIndent(createMenu(), "", "\t")
		w.WriteHeader(http.StatusOK)
		w.Write(menuAsJson)
	default:
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message": "not found"}`))
	}
}

func createMenu() Menu {
	menu := Menu{
		[]MenuElement{
			{
				Name:        "Install",
				Description: "Install projects",
			},
			{
				Name:        "Uninstall",
				Description: "Uninstall project",
			},
			{
				Name:        "Upgrade",
				Description: "Upgrade a project",
			},
			{
				Name:        "Encrypt Secrets",
				Description: "Encrypt the secrets file",
			},
			{
				Name:        "Decrypt Secrets",
				Description: "Decrypt the secrets file",
			},
			{
				Name:        "Apply Secrets",
				Description: "Apply secrets to one namespace",
			},
			{
				Name:        "Apply Secrets to all namespaces",
				Description: "Apply secrets to all namespaces",
			},
			{
				Name:        "Create Project",
				Description: "Create a project",
			},
			{
				Name:        "Create Deployment Project",
				Description: "Create a project without Jenkins to manage IP, Ingress and Loadbalancer",
			},
			{
				Name:        "Create Jenkins User Password",
				Description: "Create a password for Jenkins user",
			},
		},
	}

	return menu
}
