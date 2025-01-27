package app

import (
	"encoding/json"
	"log"
	"net/http"
)

type response struct {
	HTTPCode int `json:"status"`
	Data     any `json:"data,omitempty"`
}

func NewResponse(w http.ResponseWriter, httpCode int, data any) {
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(httpCode)

	err := json.NewEncoder(w).Encode(&response{
		HTTPCode: httpCode,
		Data:     data,
	})

	if err != nil {
		log.Printf("Error encoding response: %v", err)
		if httpCode == http.StatusOK {
			http.Error(w, "Error encoding response", http.StatusInternalServerError)
		}
	}
}

func InvalidDataResponse(w http.ResponseWriter) {
	NewResponse(w, http.StatusBadRequest, "Invalid body")
}
