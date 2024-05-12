package main

import (
	"fmt"
	"http_servers/handlers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func ok(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Received a request to '/' and printing 'ok' as the response")
	fmt.Fprintf(w, "ok\n")
}

func greeting(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Received a request to '/greeting' and printing 'hello' as the response")
	fmt.Fprintf(w, "hello\n")
}

func main() {
	log.Println("setting up HTTP handlers...")
	r := mux.NewRouter()

	r.HandleFunc("/health", handlers.GetHealthCheck).Methods("GET")
	r.HandleFunc("/greeting-message", handlers.ListMessages).Methods("GET")
	r.HandleFunc("/greeting-message", handlers.CreateMessage).Methods("POST")
	r.HandleFunc("/greeting-message/{id}", handlers.GetMessage).Methods("GET")
	r.HandleFunc("/greeting-message/{id}", handlers.UpdateMessage).Methods("PUT")
	r.HandleFunc("/greeting-message/{id}", handlers.DeleteMessage).Methods("DEL")

	log.Println("start listening port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}