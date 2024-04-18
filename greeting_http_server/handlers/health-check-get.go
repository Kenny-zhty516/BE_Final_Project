package handlers

import (
	"fmt"
	"net/http"
	"log"
)

func GetHealthCheck(w http.ResponseWriter, r *http.Request) {
	log.Println("receive health check request")
	fmt.Fprintf(w, "ok")
	log.Println("printed 'ok' to response")
}
