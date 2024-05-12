package handlers

import (
	"bytes"
	"encoding/json"
	// "http_servers/db"
	"http_servers/model"
	"net/http"
	"net/http/httptest"
	"testing"
	// "github.com/stretchr/testify/mock"

)

// type mockDBClient struct {
// 	mock.Mock
// }

// func (m *mockDBClient) ScanItems() ([]model.GreetingMessage, error) {
// 	args := m.Called()
// 	return args.Get(0).([]model.GreetingMessage), args.Error(1)
// }

// func (m *mockDBClient) PutItem(message *model.GreetingMessage) error {
// 	args := m.Called(message)
// 	return args.Error(0)
// }

// func (m *mockDBClient) GetItem(message *model.GreetingMessage) error {
// 	args := m.Called(message)
// 	return args.Error(0)
// }

// func (m *mockDBClient) UpdateItem(message *model.GreetingMessage) error {
// 	args := m.Called(message)
// 	return args.Error(0)
// }

// func (m *mockDBClient) DeleteItem(message *model.GreetingMessage) error {
// 	args := m.Called(message)
// 	return args.Error(0)
// }

// func TestUpdateMessage(t *testing.T) {
// 	// Prepare a request
// 	message := model.GreetingMessage{ID: "1", Name: "Test"}
// 	messageBytes, err := json.Marshal(message)
// 	if err != nil {
// 		t.Fatalf("Failed to marshal test message: %v", err)
// 	}

// 	req, err := http.NewRequest("POST", "/greeting-message", bytes.NewBuffer(messageBytes))
// 	if err != nil {
// 		t.Fatalf("Failed to create request: %v", err)
// 	}

// 	// Prepare a recorder
// 	rr := httptest.NewRecorder()

// 	// Call the handler function
// 	handler := http.HandlerFunc(CreateMessage)
// 	handler.ServeHTTP(rr, req)

// 	// Check the status code
// 	if status := rr.Code; status != http.StatusOK {
// 		t.Errorf("Handler returned wrong status code: got %v want %v",
// 			status, http.StatusOK)
// 	}

// 	// Check the response body
// 	expected := ``
// 	if rr.Body.String() != expected {
// 		t.Errorf("Handler returned unexpected body: got %v want %v",
// 			rr.Body.String(), expected)
// 	}

// 	// Clean up database after testing
// 	dbClient := db.NewDynamoDBClient()
// 	dbClient.DeleteItem(&message)
// }


func TestUpdateMessage(t *testing.T) {
	mockClient := &mockDBClient{}
	handler := http.HandlerFunc(UpdateMessage)
	message := model.GreetingMessage{
		ID:   "1234",
		Name: "test",
	}
	jsonBytes, err := json.Marshal(message)
	if err != nil {
		t.Fatal(err)
	}
	mockClient.On("UpdateItem", &message).Return(nil)

	req, err := http.NewRequest("PUT", "/greeting-message/1234", bytes.NewReader(jsonBytes))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}