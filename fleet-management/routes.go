package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func registerRoutes(router *httprouter.Router) {
	router.HandlerFunc(http.MethodGet, "/fleets", getFleetsHandler)
	router.HandlerFunc(http.MethodPost, "/fleets", createFleetHandler)
	router.HandlerFunc(http.MethodDelete, "/fleets/:slug", deleteFleetHandler)
	router.HandlerFunc(http.MethodPost, "/fleets/:slug/drones", addDroneToFleetHandler)
	router.HandlerFunc(http.MethodPost, "/fleets/:slug/backlog", addTaskToBacklogHandler)
}

func writeJSON(w http.ResponseWriter, data interface{}) {
	b, err := json.Marshal(data)
	if err != nil {
		log.Printf("Could not marshal data to json: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}
