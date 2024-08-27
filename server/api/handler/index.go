package handler

import (
	"log"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte(`{"message": "Hello, to start consuming this api, go to /api/rnm"}`))
	if err != nil {
		log.Fatal(err)
	}

}
