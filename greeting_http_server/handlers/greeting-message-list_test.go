package handlers

import (
	"encoding/json"
	"http_servers/model"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)


func TestListMessages(t *testing.T) {
	mockClient := &mockDBClient{}
	handler := http.HandlerFunc(ListMessages)
	message := model.GreetingMessage{
		ID:   "1234",
		Name: "test",
	}
	mockClient.On("ScanItems").Return([]model.GreetingMessage{message}, nil)

	req, err := http.NewRequest("GET", "/greeting-message", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := []model.GreetingMessage{message}
	var result []model.GreetingMessage
	err = json.NewDecoder(rr.Body).Decode(&result)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("handler returned unexpected body: got %v want %v", result, expected)
	}
}