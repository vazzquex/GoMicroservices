package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/convert", convertHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func convertHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Mehotd not allowed", http.StatusMethodNotAllowed)
		return
	}
}
