#!/bin/bash
# create-table.sh
# Script to create a DynamoDB table

# Wait for DynamoDB Local to start listening on port 8000
while ! curl -s --connect-timeout 1 http://dynamodb-local:8000 > /dev/null; do
    echo "Waiting for service on port 8000..."
    sleep 1  # Wait for 1s before checking again
done

echo "DynamoDB Local is up and running."

# AWS CLI command to create a table
aws dynamodb create-table \
    --table-name GreetingMessage \
    --attribute-definitions AttributeName=ID,AttributeType=S \
    --key-schema AttributeName=ID,KeyType=HASH \
    --provisioned-throughput ReadCapacityUnits=1,WriteCapacityUnits=1 \
    --endpoint-url http://dynamodb-local:8000

echo "Table created."
