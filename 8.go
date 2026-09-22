package main

import "fmt"

type LogEntry struct {
	IP        string 
	Code      int    
	Timestamp string 
}

func E(logs []LogEntry) []LogEntry {
	var errors []LogEntry

	for _, entry := range logs {
		if entry.Code >= 400 && entry.Code < 600 {
			errors = append(errors, entry)
		}
	}
	return errors
}

func main() {
	logs := []LogEntry{
		{"192.168.1.1", 200, "2026-09-17 10:00:00"},
		{"192.168.1.2", 404, "2026-09-17 10:01:15"},
		{"10.0.0.5", 200, "2026-09-17 10:02:30"},
		{"172.16.0.3", 500, "2026-09-17 10:03:45"},
		{"192.168.1.1", 301, "2026-09-17 10:04:00"},
		{"10.0.0.8", 403, "2026-09-17 10:05:10"},
		{"192.168.1.5", 200, "2026-09-17 10:06:20"},
		{"172.16.0.10", 502, "2026-09-17 10:07:35"},
	}

	errors := E(logs)

	fmt.Println("Записи с ошибками (4xx и 5xx):")
	for _, e := range errors {
		fmt.Printf("IP: %-15s | Код: %d | Время: %s\n", e.IP, e.Code, e.Timestamp)
	}
	fmt.Println("\nВсего ошибок:", len(errors))
}