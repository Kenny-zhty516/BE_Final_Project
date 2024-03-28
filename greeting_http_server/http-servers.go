package main

import (
	"fmt"
	"net/http"
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
	http.HandleFunc("/", ok)
	http.HandleFunc("/greeting", greeting)

	http.ListenAndServe(":8080", nil)
}
