package handlers

import (
	"encoding/json"
	"http_servers/model"
	"log"
	"net/http"
)

func CreateMessage(w http.ResponseWriter, r *http.Request) {
	log.Println("receive create messages request")

	// Unmarshal
	var message model.GreetingMessage
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&message)
	if err != nil {
		log.Printf("Unmarshal request to JSON failed with error: %v", err)
		http.Error(w, "Failed to decode request body to JSON", http.StatusInternalServerError)
		return
	}

	err = dbClient.PutItem(&message)
	if err != nil {
		log.Printf("save message failed with error: %v", err)
		http.Error(w, "Failed to save message", http.StatusInternalServerError)
		return
	}
	log.Printf("created message with id: %v, name: %v", message.ID, message.Name)
}
