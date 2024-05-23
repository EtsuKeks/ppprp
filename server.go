package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var timeRequestsCount int

func getCurrentTime(w http.ResponseWriter, r *http.Request) {
	timeRequestsCount++

	resp, err := http.Get("http://worldtimeapi.org/api/timezone/Europe/Moscow")
	if err != nil {
		fmt.Printf("Error")
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Error")
		return
	}

	var result map[string]string
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		fmt.Printf("Error")
		return
	}

	time, ok := result["datetime"]
	if !ok {
		fmt.Println("Error")
		return
	}

	response := map[string]string{"current_time": time}
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
