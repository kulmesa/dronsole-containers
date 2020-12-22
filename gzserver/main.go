package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	port := "8081"

	router := httprouter.New()
	registerRoutes(router)

	log.Printf("Listening on port %s", port)
	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		log.Fatal(err)
	}
}
