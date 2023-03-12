package main

import (
	"hapoon/api/handler"
	"log"
	"net/http"
)

const (
	name    = "hapoon-api"
	version = "0.1.0"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/status", handler.Status)

	log.Fatal(http.ListenAndServe(":8080", mux))
}
