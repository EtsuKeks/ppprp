package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

var timeRequestsCount int

func getCurrentTime(w http.ResponseWriter, r *http.Request) {
	timeRequestsCount++
	currentTime := time.Now().Format(time.UnixDate)
	response := map[string]string{"current_time": currentTime}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

func getTimeRequestsStatistics(w http.ResponseWriter, r *http.Request) {
	response := map[string]int{"time_requests_count": timeRequestsCount}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

func main() {
	http.HandleFunc("/time", getCurrentTime)
	http.HandleFunc("/statistics", getTimeRequestsStatistics)
	if err := http.ListenAndServe(":5000", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
