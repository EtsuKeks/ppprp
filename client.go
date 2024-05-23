package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

func fetchStatisticsAndWriteToFile() {
	resp, err := http.Get("http://service/statistics")
	if err != nil {
		fmt.Printf("Error")
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Error")
		return
	}

	var result map[string]int
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		fmt.Printf("Error")
		return
	}

	timeRequestsCount, ok := result["time_requests_count"]
	if !ok {
		fmt.Println("Error")
		return
	}

	file, err := os.OpenFile("statistics.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("Error")
		return
	}
	defer file.Close()

	if _, err := file.WriteString(fmt.Sprintf("%d\n", timeRequestsCount)); err != nil {
		fmt.Printf("Error")
	}
}

func main() {
	for {
		fetchStatisticsAndWriteToFile()
		time.Sleep(5 * time.Second)
	}
}
