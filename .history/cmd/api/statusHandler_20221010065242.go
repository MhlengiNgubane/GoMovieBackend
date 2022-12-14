package main

import (
	"log"
	"net/http"
)

func (app *application) statusHandler(w http.ResponseWriter, r *http.Request) {
	currentStatus := AppStatus {
		Status: "Available",
		Environment: cfg.env,
		Version: version,
	}

	js, err := json.MarshalIndent(currentStatus, "", "\t")
	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(js)
}