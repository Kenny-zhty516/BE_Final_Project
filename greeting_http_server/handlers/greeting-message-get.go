package handlers

import (
	"net/http"
	"context"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func GetMessage(w http.ResponseWriter, r *http.Request) {
	// Implementation of Get item from DynamoDB


	Id := r.URL.Query().Get("ID")
	Name := r.URL.Query().Get("Name")

	tableName := os.Getenv("DYNAMODB_TABLE_NAME")

	input := &dynamodb.GetItemInput{
		TableName: aws.String(tableName),
		Key: map[string]types.AttributeValue{
			"ID":   &types.AttributeValueMemberS{Value: Id},
			"Name": &types.AttributeValueMemberS{Value: Name},
		},
	}

	_, err := dbClient.SVC.GetItem(context.TODO(), input)
	if err != nil {
		http.Error(w, "Failed to put item to DynamoDB", http.StatusInternalServerError)
		return
	}
}
