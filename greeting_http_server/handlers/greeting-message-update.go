package handlers

import (
	"encoding/json"
	"http_servers/model"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func UpdateMessage(w http.ResponseWriter, r *http.Request) {
	log.Println("received update messages request")


	// Extract ID from URL path parameters
	vars := mux.Vars(r)
	id := vars["id"]

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

	err = dbClient.UpdateItem(id, message.Name)
	if err != nil {
		log.Printf("update message failed with error: %v", err)
		http.Error(w, "Failed to updated message", http.StatusInternalServerError)
		return
	}
	// log.Printf("updated message with id: %v, name: %v", message.ID, message.Name)
	log.Printf("updated message with id: %v", id)
	
}
