package handlers

import (
	// "context"
	"http_servers/model"
	// "json"
	"log"
	"net/http"
	"encoding/json"
	// "os"

	// "github.com/aws/aws-sdk-go-v2/aws"
	// "github.com/aws/aws-sdk-go-v2/"
	// "github.com/aws/aws-sdk-go-v2/service/dynamodb"
	// "github.com/aws/aws-sdk-go-v2/service/dynamodb/dynamodbattribute"
	// "github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	// "github.com/aws/aws-sdk-go-v2/session"
)

func GetMessage(w http.ResponseWriter, r *http.Request) {
	// Implementation of Get item from DynamoDB

	log.Println("received get messages request")

	var message model.GreetingMessage
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&message)
	if err != nil {
		log.Printf("Unmarshal request to JSON failed with error: %v", err)
		http.Error(w, "Failed to decode request body to JSON", http.StatusInternalServerError)
		return
	}

	err = dbClient.GetItem(&message)
	if err != nil {
		log.Printf("get message failed with error: %v", err)
		http.Error(w, "Failed to get message", http.StatusInternalServerError)
		return
	}
	log.Printf("retrieved message with id: %v, name: %v", message.ID, message.Name)

	// // Get the item ID from the request parameters
	// ID := r.URL.Query().Get("ID")
	// Name := r.URL.Query().Get("Name")

	// // Check if the item exists
	// if ID == "" {
	// 	log.Printf("item with ID %s not found", ID)
	// 	http.Error(w, "Item not found", http.StatusNotFound)
	// 	return
	// }

	// if Name == "" {
	// 	log.Printf("item with Name %s not found", Name)
	// 	http.Error(w, "Item not found", http.StatusNotFound)
	// 	return
	// }

	// tableName := os.Getenv("DYNAMODB_TABLE_NAME")
	// input := &dynamodb.GetItemInput{
	// 	TableName: aws.String(tableName),
	// 	Key: map[string]types.AttributeValue{
	// 		"ID":   &types.AttributeValueMemberS{Value: ID},
	// 		"Name": &types.AttributeValueMemberS{Value: Name},
	// 	},
	// }

	// Retrieve the item from DynamoDB
	// item, err := dbClient.SVC.GetItem(context.TODO(), id)
	// _, err := dbClient.SVC.GetItem(context.TODO(), input)

	// if err != nil {

	// 	log.Printf("failed to retrieve item with ID %s and name %s: %v", ID, Name, err)
	// 	http.Error(w, "Failed to retrieve item", http.StatusInternalServerError)
	// 	return
	// }

	// return

}
