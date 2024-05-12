package db

import (
	"os"
	"fmt"
	"log"
	"time"
	"context"
	"http_servers/model"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
)

type DynamoDBClient struct {
	svc *dynamodb.Client
}

type DynamoDBAPI interface {
	ScanItems() ([]model.GreetingMessage, error)
	PutItem(message *model.GreetingMessage) error
	GetItem(message *model.GreetingMessage) error
	UpdateItem(message *model.GreetingMessage) error
	DeleteItem(message *model.GreetingMessage) error
	QueryByID(id string) (string, error)
}

func NewDynamoDBClient() *DynamoDBClient {
	dynamodb_local_endpoint := os.Getenv("DYNAMODB_LOCAL_ENDPOINT")
	customResolver := aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
		if service == dynamodb.ServiceID && dynamodb_local_endpoint != "" {
			return aws.Endpoint{
				URL: dynamodb_local_endpoint,
			}, nil
		}
		
		return aws.Endpoint{}, &aws.EndpointNotFoundError{}
	})

	sdkConfig, err := config.LoadDefaultConfig(context.TODO(), config.WithEndpointResolverWithOptions(customResolver))
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

	svc := dynamodb.NewFromConfig(sdkConfig)

	return &DynamoDBClient{svc: svc}
}

func (d *DynamoDBClient) ScanItems() ([]model.GreetingMessage, error) {
	var items []model.GreetingMessage

	tableName := os.Getenv("DYNAMODB_TABLE_NAME")

	data, err := d.svc.Scan(context.TODO(), &dynamodb.ScanInput{
		TableName: aws.String(tableName),
	})
	if err != nil {
		return items, fmt.Errorf("Query: %v\n", err)
	}

	err = attributevalue.UnmarshalListOfMaps(data.Items, &items)
	if err != nil {
		return items, fmt.Errorf("UnmarshalListOfMaps: %v\n", err)
	}

	return items, nil
}

func (d *DynamoDBClient) PutItem(message *model.GreetingMessage) error {
	tableName := os.Getenv("DYNAMODB_TABLE_NAME")
	currentTime := time.Now().Format(time.RFC3339) 

	input := &dynamodb.PutItemInput{
		TableName: aws.String(tableName),
		Item: map[string]types.AttributeValue{
			"ID":         &types.AttributeValueMemberS{Value: message.ID},
			"Name":       &types.AttributeValueMemberS{Value: message.Name},
			"created_at": &types.AttributeValueMemberS{Value: currentTime},
		},
	}
	if _, err := d.svc.PutItem(context.TODO(), input); err != nil {
		return fmt.Errorf("failed to put item (ID: %v, Name: %v) into DynamoDB: %w", message.ID, message.Name, err)
	}
	return nil
}

func (d *DynamoDBClient) GetItem(message *model.GreetingMessage) error {
	tableName := os.Getenv("DYNAMODB_TABLE_NAME")

	input := &dynamodb.GetItemInput{
		TableName: aws.String(tableName),
		Key: map[string]types.AttributeValue{
			"ID":   &types.AttributeValueMemberS{Value: message.ID},
			"Name": &types.AttributeValueMemberS{Value: message.Name},
		},
	}

	if item, err := d.svc.GetItem(context.TODO(), input); err != nil {
		return fmt.Errorf("failed to get item (ID: %v, Name: %v) from DynamoDB: %w", message.ID, message.Name, err)
	} else {
		fmt.Print(item)
		return nil
	}
}

func (d *DynamoDBClient) UpdateItem(id string, newName string) error {
	tableName := os.Getenv("DYNAMODB_TABLE_NAME")
	currentTime := time.Now().Format(time.RFC3339) 


	input := &dynamodb.UpdateItemInput{
		TableName: aws.String(tableName),
		Key: map[string]types.AttributeValue{
			"ID": &types.AttributeValueMemberS{Value: id},
		},
		UpdateExpression: aws.String("SET #N = :newName, #U = :currentTime"), 
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":newName": &types.AttributeValueMemberS{Value: newName},
			":currentTime": &types.AttributeValueMemberS{Value: currentTime},
		},
		ExpressionAttributeNames: map[string]string{
			"#N": "Name", 
			"#U": "updated_at", 
		},
	}

	if item, err := d.svc.UpdateItem(context.TODO(), input); err != nil {
		return fmt.Errorf("failed to update item (ID: %v, Name: %v) from DynamoDB: %w", id, newName, err)
	} else {
		fmt.Print(item)
		return nil
	}
}

func (d *DynamoDBClient) DeleteItem(id string) error {
	tableName := os.Getenv("DYNAMODB_TABLE_NAME")

	input := &dynamodb.DeleteItemInput{
		TableName: aws.String(tableName),
		Key: map[string]types.AttributeValue{
			"ID": &types.AttributeValueMemberS{Value: id},
		},
	}

	if item, err := d.svc.DeleteItem(context.TODO(), input); err != nil {
		return fmt.Errorf("failed to delete item (ID: %v, Name: %v) from DynamoDB: %w", id, err)
	} else {
		fmt.Print(item)
		return nil
	}
}

func (d *DynamoDBClient) QueryByID(id string) (string, error) {
	tableName := os.Getenv("DYNAMODB_TABLE_NAME")

	input := &dynamodb.QueryInput{
		TableName:              aws.String(tableName),
		KeyConditionExpression: aws.String("ID = :id"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":id": &types.AttributeValueMemberS{Value: id},
		},
		ProjectionExpression: aws.String("#name"), 
		ExpressionAttributeNames: map[string]string{
			"#name": "Name", 
		},
	}

	result, err := d.svc.Query(context.TODO(), input)
	if err != nil {
		return "", fmt.Errorf("failed to query item by ID: %w", err)
	}

	if len(result.Items) == 0 {
		return "", fmt.Errorf("no item found with ID: %s", id)
	}

	item := result.Items[0]
	nameAttr, found := item["Name"]
	if !found {
		return "", fmt.Errorf("name attribute not found for ID: %s", id)
	}

	name, ok := nameAttr.(*types.AttributeValueMemberS)
	if !ok {
		return "", fmt.Errorf("unexpected name attribute type for ID: %s", id)
	}

	return name.Value, nil
}
