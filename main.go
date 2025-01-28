package main

import (
	"log"
	"net/http"

	api "github.com/hapoon/api/gen"
)

const (
	name    = "hapoon-api"
	version = "0.1.0"
)

func main() {
	log.Println("Server started")
	StatusAPIService := api.NewStatusAPIService()
	StatusAPIController := api.NewStatusAPIController(StatusAPIService)

	router := api.NewRouter(StatusAPIController)

	log.Fatal(http.ListenAndServe(":8080", router))
}
