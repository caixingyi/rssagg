package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func responseWithJSON(w http.ResponseWriter, code int, payload any) {
	data, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Failed to marshal JSON response: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(data)
}

func responseWithError(w http.ResponseWriter, code int, msg string) {
	if code > 499 {
		log.Println(msg)
	}
	responseWithJSON(w, code, map[string]string{"error": msg})
}
