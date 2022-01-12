package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"shs/config"
	"shs/handlers"
	"shs/router"
	"shs/tools"
)

const portNumber = ":8080"

// main is the main function
func main() {

	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	handlers.NewRepo(&app)

	render.NewTemplates(&app)

	rh := mux.NewRouter()

	router.HTTPRouter(rh)

	fmt.Println(fmt.Sprintf("Staring application on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, rh)

}
