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