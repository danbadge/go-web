package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Response struct {
	Message string ` json:"message" `
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/ping", PingHandler)
	log.Fatal(http.ListenAndServe(":8002", r))
}

func PingHandler(w http.ResponseWriter, _r *http.Request) {
	payload := Response{Message: "pong"}
	jsonResponse, jsonError := json.Marshal(payload)

	if jsonError != nil {
		fmt.Println("Unable to encode JSON")
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}
