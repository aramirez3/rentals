package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithJSON(w http.ResponseWriter, statusCode int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	responseData, err := json.Marshal(payload)

	if err != nil {
		log.Printf("Error marshalling json: %s\n", err)
		w.WriteHeader(500)
		return
	}
	w.WriteHeader(statusCode)
	_, err = w.Write(responseData)
	if err != nil {
		log.Printf("Error sending response json: %s\n", err)
	}
}
