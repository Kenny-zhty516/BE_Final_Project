package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type Item struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func main() {
	lambda.Start(handleRequest)
}

func handleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	item, err := parseRequestBody(request.Body)
	if err != nil {
		log.Printf("Error parsing request body: %v", err)
		return clientError(400, err.Error())
	}

	if err := putItem(ctx, item); err != nil {
		log.Printf("Error putting item to DynamoDB: %v", err)
		return serverError(err)
	}

	log.Printf("Successfully added item: %v", item)
	return successResponse(fmt.Sprintf("Item with ID %s added.", item.ID))
}

func parseRequestBody(body string) (*Item, error) {
	var item Item
	if err := json.Unmarshal([]byte(body), &item); err != nil {
		return nil, err
	}
	return &item, nil
}

func putItem(ctx context.Context, item *Item) error {
	svc, err := getDynamoDBClient(ctx)
	if err != nil {
		return fmt.Errorf("failed to get DynamoDB client: %w", err)
	}
	tableName := os.Getenv("DYNAMODB_TABLE_NAME")
	log.Printf("DYNAMODB_TABLE_NAME: %v", tableName)
	input := &dynamodb.PutItemInput{
		TableName: aws.String(tableName),
		Item: map[string]types.AttributeValue{
			"ID":   &types.AttributeValueMemberS{Value: item.ID},
			"Name": &types.AttributeValueMemberS{Value: item.Name},
		},
	}

	if _, err := svc.PutItem(ctx, input); err != nil {
		return fmt.Errorf("failed to put item into DynamoDB: %w", err)
	}
	return nil
}

func getDynamoDBClient(ctx context.Context) (*dynamodb.Client, error) {
	// Load AWS configuration with custom DynamoDB endpoint if specified
	dynamodb_local_endpoint := os.Getenv("DYNAMODB_LOCAL_ENDPOINT")
	log.Printf("DYNAMODB_LOCAL_ENDPOINT: %v", dynamodb_local_endpoint)
	customResolver := aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
		if service == dynamodb.ServiceID && dynamodb_local_endpoint != "" {
			return aws.Endpoint{
				URL: dynamodb_local_endpoint,
			}, nil
		}
		// returning EndpointNotFoundError will allow the service to fallback to it's default resolution
		return aws.Endpoint{}, &aws.EndpointNotFoundError{}
	})

	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithEndpointResolverWithOptions(customResolver))

	if err != nil {
		return nil, err
	}
	return dynamodb.NewFromConfig(cfg), nil
}

func clientError(status int, msg string) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{StatusCode: status, Body: msg}, nil
}

func serverError(err error) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{StatusCode: 500, Body: err.Error()}, nil
}

func successResponse(message string) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{StatusCode: 200, Body: message}, nil
}
