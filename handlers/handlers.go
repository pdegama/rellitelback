package handlers

import (
	"shs/config"
)

// Repo the repository used by the handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) {
	Repo = &Repository{
		App: a,
	}
}
