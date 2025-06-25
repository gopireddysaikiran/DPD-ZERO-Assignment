package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		response := map[string]string{
			"status":  "ok",
			"service": "1",
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	})

	log.Println("Service 1 listening on port 8001...")
	log.Fatal(http.ListenAndServe("0.0.0.0:8001", nil))
}
