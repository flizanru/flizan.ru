package main

import (
	"net"
	"net/http"
	"log"
	"time"
	"encoding/json"
	"sync"
)

type ServerStatus struct {
	IP     string  `json:"ip"`
	Status string  `json:"status"`
	Ping   float64 `json:"ping"` 
}

var serverStatuses = []ServerStatus{
	{IP: "46.174.48.244", Status: "offline", Ping: 0},
	{IP: "89.39.121.159", Status: "offline", Ping: 0},
	{IP: "87.120.166.96", Status: "offline", Ping: 0},
}

var statusesMutex = sync.Mutex{}

func pingServer(ip string) (string, float64) {
	start := time.Now()
	_, err := net.DialTimeout("tcp", ip+":22", 5*time.Second)
	duration := time.Since(start)

	if err != nil {
		return "offline", 0
	}
	return "online", duration.Seconds() * 1000
}

func updateServerStatuses() {
	for {
		statusesMutex.Lock()
		for i := range serverStatuses {
			status, ping := pingServer(serverStatuses[i].IP)
			serverStatuses[i].Status = status
			serverStatuses[i].Ping = ping
		}
		statusesMutex.Unlock()

		time.Sleep(1 * time.Minute)
	}
}

func serverStatusHandler(w http.ResponseWriter, r *http.Request) {
	statusesMutex.Lock()
	defer statusesMutex.Unlock()

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(serverStatuses)
	if err != nil {
		http.Error(w, "Error encoding JSON response", http.StatusInternalServerError)
	}
}

func main() {
	go updateServerStatuses()

	http.HandleFunc("/status", serverStatusHandler)

	log.Println("Server started on port 2008...")
	err := http.ListenAndServe(":2008", nil)
	if err != nil {
		log.Fatal("Server error:", err)
	}
}
