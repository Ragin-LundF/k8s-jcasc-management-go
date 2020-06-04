package server

import (
	"encoding/json"
	"k8s-management-go/app/cli/menu"
	"net/http"
)

type Menu struct {
	Elements []menu.MenuitemModel
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
	menuitemsStructure := Menu{menu.CreateMenuItems()}

	return menuitemsStructure
}
