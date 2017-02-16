package server

import (
	"net/http"
	"github.com/gorilla/mux"
	sw "github.com/predixdeveloperACN/swagger-ui"
	"os"
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
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	http.ListenAndServe(":"+port, routes)
}
