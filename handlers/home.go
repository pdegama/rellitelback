package handlers

import (
	"net/http"
	"shs/models"
	"shs/tools"
)

// Home is the handlers for the home page
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}
