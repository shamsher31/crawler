package crawler

import (
	"github.com/gorilla/mux"
	"github.com/shamsher31/crawler/middleware"
)

// Router register necessary routes and returns an instance of a router.
func Router() *mux.Router {
	router := mux.NewRouter()
	router.Use(middleware.Headers)

	router.HandleFunc("/health-check", HealthCheck).Methods("GET")
	router.HandleFunc("/", Run).Methods("GET")

	return router
}
