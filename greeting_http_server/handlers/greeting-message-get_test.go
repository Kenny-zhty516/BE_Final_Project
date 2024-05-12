package handlers

import (
	"bytes"
	"encoding/json"
	// "errors"
	"http_servers/model"
	"net/http"
	"net/http/httptest"
	// "reflect"
	"testing"

	"github.com/stretchr/testify/mock"
)

type mockDBClient struct {
	mock.Mock
}

func (m *mockDBClient) ScanItems() ([]model.GreetingMessage, error) {
	args := m.Called()
	return args.Get(0).([]model.GreetingMessage), args.Error(1)
}

func (m *mockDBClient) PutItem(message *model.GreetingMessage) error {
	args := m.Called(message)
	return args.Error(0)
}

func (m *mockDBClient) GetItem(message *model.GreetingMessage) error {
	args := m.Called(message)
	return args.Error(0)
}

func (m *mockDBClient) UpdateItem(message *model.GreetingMessage) error {
	args := m.Called(message)
	return args.Error(0)
}

func (m *mockDBClient) DeleteItem(message *model.GreetingMessage) error {
	args := m.Called(message)
	return args.Error(0)
}

func TestGetMessage(t *testing.T) {
	mockClient := &mockDBClient{}
	handler := http.HandlerFunc(GetMessage)
	message := model.GreetingMessage{
		ID:   "1234",
		Name: "test",
	}
	jsonBytes, err := json.Marshal(message)
	if err != nil {
		t.Fatal(err)
	}
	mockClient.On("GetItem", &message).Return(nil)

	req, err := http.NewRequest("GET", "/greeting-message/1234", bytes.NewReader(jsonBytes))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}
