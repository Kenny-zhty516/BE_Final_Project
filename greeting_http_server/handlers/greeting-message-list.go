package handlers

import (
	"encoding/json"
	"http_servers/db"
	"log"
	"net/http"
)

var dbClient = db.NewDynamoDBClient()

func ListMessages(w http.ResponseWriter, r *http.Request) {
	log.Println("receive list messages request")
	items, err := dbClient.ScanItems()
	if err != nil {
		log.Printf("retrieve failed with error: %v", err)
		http.Error(w, "Failed to retrieve items", http.StatusInternalServerError)
		return
	}

	log.Printf("retrieved %v items", len(items))
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(items); err != nil {
		log.Printf("encode failed with error: %v", err)
		http.Error(w, "Failed to encode items", http.StatusInternalServerError)
		return
	}
	log.Printf("returned all items")
}
