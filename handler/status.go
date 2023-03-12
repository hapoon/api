package handler

import (
	"encoding/json"
	"log"
	"net/http"
)

type StatusResponse struct {
	State string `json:"state"`
}

func Status(w http.ResponseWriter, r *http.Request) {
	s := StatusResponse{State: "success"}
	res, err := json.Marshal(s)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}
