package db

import (
	"context"
	"fmt"
	"http_servers/model"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
)

type DynamoDBClient struct {
	SVC *dynamodb.Client
}

func NewDynamoDBClient() *DynamoDBClient {
	dynamodb_local_endpoint := "http://dynamodb-local:8000"
	customResolver := aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
		if service == dynamodb.ServiceID && dynamodb_local_endpoint != "" {
			return aws.Endpoint{
				URL: dynamodb_local_endpoint,
			}, nil
		}
		// returning EndpointNotFoundError will allow the service to fallback to it's default resolution
		return aws.Endpoint{}, &aws.EndpointNotFoundError{}
	})

	sdkConfig, err := config.LoadDefaultConfig(context.TODO(), config.WithEndpointResolverWithOptions(customResolver))
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

	SVC := dynamodb.NewFromConfig(sdkConfig)

	return &DynamoDBClient{SVC: SVC}
}

func (d *DynamoDBClient) ScanItems() ([]model.GreetingMessage, error) {
	var items []model.GreetingMessage

	data, err := d.SVC.Scan(context.TODO(), &dynamodb.ScanInput{
		TableName: aws.String("GreetingMessage"),
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
