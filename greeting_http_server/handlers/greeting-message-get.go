package handlers

import (
	"encoding/json"
	"http_servers/model"
	"log"
	"net/http"
)

func GetMessage(w http.ResponseWriter, r *http.Request) {
	log.Println("received get messages request")

	var message model.GreetingMessage
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&message)
	if err != nil {
		log.Printf("Unmarshal request to JSON failed with error: %v", err)
		http.Error(w, "Failed to decode request body to JSON", http.StatusInternalServerError)
		return
	}

	err = dbClient.GetItem(&message)
	if err != nil {
		log.Printf("get message failed with error: %v", err)
		http.Error(w, "Failed to get message", http.StatusInternalServerError)
		return
	}
	log.Printf("retrieved message with id: %v, name: %v", message.ID, message.Name)

}
