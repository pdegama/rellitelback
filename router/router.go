package router

import (
	"github.com/gorilla/mux"
	"net/http"
	"shs/handlers"
)

func HTTPRouter(rh *mux.Router) {

	rh.HandleFunc("/home", handlers.Repo.Home)
	rh.HandleFunc("/about/{name}", handlers.Repo.About)

	rh.NotFoundHandler = http.HandlerFunc(handlers.Repo.Home)

}
