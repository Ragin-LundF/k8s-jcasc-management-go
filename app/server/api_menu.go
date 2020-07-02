package server

import (
	"encoding/json"
	"k8s-management-go/app/cli/menu"
	"net/http"
)

// Menu defines the structure of an array of menu items
type Menu struct {
	Elements []menu.MenuitemModel
}

// MenuAPI is the API for the menu values
func MenuAPI(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case "GET":
		menuAsJSON, _ := json.MarshalIndent(createMenu(), "", "\t")
		w.WriteHeader(http.StatusOK)
		w.Write(menuAsJSON)
	default:
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message": "not found"}`))
	}
}

func createMenu() Menu {
	menuitemsStructure := Menu{menu.CreateMenuItems()}

	return menuitemsStructure
}
