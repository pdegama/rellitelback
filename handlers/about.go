package handlers

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"shs/models"
	"shs/tools"
)

// About is the handlers for the about page
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	log.Println(vars["name"])
	// perform some logic
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again"

	// send data to the template
	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}
