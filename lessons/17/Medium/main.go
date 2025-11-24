package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type PowerData struct {
	PowerGenerated int    `json:"power_generated"`
	Units          string `json:"units"`
}

func main() {

	http.HandleFunc("/", rootHandler)

	http.HandleFunc("/status", statusHandler)

	http.HandleFunc("/data", dataHandler)

	fmt.Println("Server deployed at 8081...")

	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		log.Fatal("Error at server deployment:", err)
	}
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Server is running")
}

func statusHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Status: OK")
}

func dataHandler(w http.ResponseWriter, r *http.Request) {

	data := PowerData{
		PowerGenerated: 1500,
		Units:          "watts",
	}

	w.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		http.Error(w, "Error generating JSON", http.StatusInternalServerError)
	}
}
