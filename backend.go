package main

import (
	"net"
	"net/http"
	"time"
	"log"
	"encoding/json"
)

type ServerStatus struct {
	IP     string  `json:"ip"`
	Status string  `json:"status"`
	Ping   float64 `json:"ping"` 
}

func pingServer(ip string) (string, float64) {
	start := time.Now()
	_, err := net.DialTimeout("tcp", ip+":22", 5*time.Second)
	duration := time.Since(start)

	if err != nil {
		return "offline", 0
	}
	return "online", duration.Seconds() * 1000 
}

func serverStatusHandler(w http.ResponseWriter, r *http.Request) {
	servers := []string{
		"46.174.48.244", 
		"89.39.121.159",
		"87.120.166.96",  
	}

	var statuses []ServerStatus

	for _, ip := range servers {
		status, ping := pingServer(ip)
		statuses = append(statuses, ServerStatus{
			IP:     ip,
			Status: status,
			Ping:   ping,
		})
	}

	w.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(w).Encode(statuses)
	if err != nil {
		http.Error(w, "Error encoding JSON response", http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/status", serverStatusHandler)
	log.Println("Server started on port 2008...")
	err := http.ListenAndServe(":2008", nil)
	if err != nil {
		log.Fatal("Server error:", err)
	}
}
