package handlers

import (
	"database/sql"
	"net/http"

	"github.com/gorilla/securecookie"
)

// Repo the repository used by the handlers
var Repo = &Repository{}

// Repository is the repository type
type Repository struct {
	Headers      *map[string]string
	DataBase     *sql.DB
	SecureCookie *securecookie.SecureCookie
	StaticURL    string
	StaticURL2   string
	FrontURL     string
	SecKey       string
}

func SetHeaders(m *Repository, w http.ResponseWriter) {
	for key, Value := range *m.Headers {
		w.Header().Set(key, Value)
	}
}
