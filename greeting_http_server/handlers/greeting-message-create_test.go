package handlers

import (
	"bytes"
	"encoding/json"
	"http_servers/db"
	"http_servers/model"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateMessage(t *testing.T) {
	// Prepare a request
	message := model.GreetingMessage{ID: "1234", Name: "Test"}
	messageBytes, err := json.Marshal(message)
	if err != nil {
		t.Fatalf("Failed to marshal test message: %v", err)
	}

	req, err := http.NewRequest("POST", "/greeting-message", bytes.NewBuffer(messageBytes))
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	// Prepare a recorder
	rr := httptest.NewRecorder()

	// Call the handler function
	handler := http.HandlerFunc(CreateMessage)
	handler.ServeHTTP(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body
	expected := ``
	if rr.Body.String() != expected {
		t.Errorf("Handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}

	// Clean up database after testing
	dbClient := db.NewDynamoDBClient()
	dbClient.DeleteItem(message.ID)
}
