package main

import (
	"http_servers/handlers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	log.Println("setting up HTTP handlers...")
	r := mux.NewRouter()

	r.HandleFunc("/health", handlers.GetHealthCheck).Methods("GET")
	r.HandleFunc("/greeting-message", handlers.ListMessages).Methods("GET")
	r.HandleFunc("/greeting-message", handlers.CreateMessage).Methods("POST")
	r.HandleFunc("/greeting-message/{id}", handlers.GetMessage).Methods("GET")
	r.HandleFunc("/greeting-message/{id}", handlers.UpdateMessage).Methods("PUT")

	log.Println("start listening port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
