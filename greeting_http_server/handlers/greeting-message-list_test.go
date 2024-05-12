package handlers

import (
	"encoding/json"
	"github.com/stretchr/testify/mock"
	"http_servers/model"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

// Mocking DynamoDB client
type MockDynamoDBClient struct {
	mock.Mock
}

func (m *MockDynamoDBClient) ScanItems() ([]model.GreetingMessage, error) {
	args := m.Called()
	return args.Get(0).([]model.GreetingMessage), args.Error(1)
}

func TestListMessages(t *testing.T) {
	// Mocking DynamoDB client
	mockDBClient := new(MockDynamoDBClient)

	// Expected response from DynamoDB client
	expectedMessages := []model.GreetingMessage{
		{ID: "1", Name: "Message 1"},
		{ID: "2", Name: "Message 2"},
	}

	// Configure mock behavior
	mockDBClient.On("ScanItems").Return(expectedMessages, nil)

	// Inject mockDBClient into handler
	dbClient = mockDBClient

	// Create a request to pass to our handler
	req, err := http.NewRequest("GET", "/greeting-message", nil)
	if err != nil {
		t.Fatal(err)
	}

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response
	rr := httptest.NewRecorder()

	// Call the handler function
	handler := http.HandlerFunc(ListMessages)
	handler.ServeHTTP(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body
	var messages []model.GreetingMessage
	if err := json.Unmarshal(rr.Body.Bytes(), &messages); err != nil {
		t.Errorf("error unmarshalling response body: %v", err)
	}

	if !reflect.DeepEqual(messages, expectedMessages) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			messages, expectedMessages)
	}

	// Assert that the ScanItems method was called once
	mockDBClient.AssertCalled(t, "ScanItems")
}
