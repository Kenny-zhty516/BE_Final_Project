package handlers

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func DeleteMessage(w http.ResponseWriter, r *http.Request) {
	log.Println("received delete messages request")

	// Extract ID from URL path parameters
	vars := mux.Vars(r)
	id := vars["id"]

	// Delete the item using the retrieved ID and name
	err := dbClient.DeleteItem(id)
	if err != nil {
		log.Printf("delete message failed with error: %v", err)
		http.Error(w, "Failed to delete message", http.StatusInternalServerError)
		return
	}
	log.Printf("deleted message with id: %v", id)

}
