package server

import (
	"net/http"
	"github.com/gorilla/mux"
	sw "github.com/predixdeveloperACN/swagger-ui"
)

var routes *mux.Router

func SetupServer() {
	routes = mux.NewRouter()

	// simple route
	routes.HandleFunc("/hello", handleHello)

	// file server
	routes.Handle("/", http.FileServer(http.Dir("./static")))

	// rest methods
	routes.HandleFunc("/languages", handleGetLanguages).Methods("GET")
	routes.HandleFunc("/languages/{language}", handleGetLanguage).Methods("GET")
	routes.HandleFunc("/language", handlePostLanguage).Methods("POST")

	// attach the swagger routes
	sw.AttachSwaggerUI(routes, "/")
}

func StartServer() {
	http.ListenAndServe(":8080", routes)
}
