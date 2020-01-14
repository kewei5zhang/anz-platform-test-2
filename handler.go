package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Handler struct {
	Router *mux.Router
}

func (h *Handler) Initialize() {
	h.Router = mux.NewRouter().StrictSlash(true)
	h.initializeRoutes()
}

func (h *Handler) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, h.Router))
}

func (h *Handler) initializeRoutes() {
	initApplications()
	h.Router.HandleFunc("/", homeLink)
	h.Router.HandleFunc("/version", getAppVersion).Methods("GET")
}
