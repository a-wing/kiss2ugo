package api

import (
	"net/http"

	"kiss2u/kiss"
	"kiss2u/storage"

	"github.com/gorilla/mux"
)

// Serve declares API routes for the application.
func Serve(router *mux.Router, store *storage.Storage, kiss *kiss.Kiss) {
	handler := &handler{store, kiss}

	sr := router.PathPrefix("/api/v2").Subrouter()
	sr.Use(handleCORS)
	sr.Methods(http.MethodOptions)

	sr.HandleFunc("/test", handler.test).Methods(http.MethodGet)
	sr.PathPrefix("/docs/").Handler(http.StripPrefix("/api/v2/docs/", http.FileServer(http.Dir("docs/")))).Methods(http.MethodGet)

	sr.HandleFunc("/packages", handler.pkgs).Methods(http.MethodGet)
	sr.HandleFunc("/packages/{name}", handler.findPkg).Methods(http.MethodGet)
	sr.HandleFunc("/packages/{name}/logs/{timestamp:[0-9]+}", handler.findLog).Methods(http.MethodGet)

	sr.HandleFunc("/users", handler.users).Methods(http.MethodGet)
	sr.HandleFunc("/users/{name}", handler.findUserPkg).Methods(http.MethodGet)

	sr.HandleFunc("/webhooks/sync", handler.hookSync).Methods(http.MethodPost)
}
