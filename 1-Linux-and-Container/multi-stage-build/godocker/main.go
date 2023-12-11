package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Time struct {
	CurrentTime string `json:"current_time"`
}

func main() {
	// defining router
	mux := http.NewServeMux()
	mux.HandleFunc("/time", getTime)

	// starting server
	log.Fatal(http.ListenAndServe( "0.0.0.0:3000", mux))
}

func getTime(w http.ResponseWriter, r *http.Request) {
	currentTime := []Time{
		{ CurrentTime: http.TimeFormat },
	}

	json.NewEncoder(w).Encode(currentTime)
}
