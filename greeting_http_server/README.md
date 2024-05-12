# CDK Python Project Guide

Welcome to your CDK Python project! This guide will walk you through the steps to test, create, update, list, and delete items using the HTTP server in your local development environment.


## Setting Up Your Local Environment

To set up your local development and testing environment, we'll be using Docker Compose. Navigate to the `local_dev` directory and run the `make` command as shown below:

cd local_dev
make

## Running Tests

To run tests, you'll need to enter the Docker container. Navigate to the `greeting_http_server/handlers` directory and execute the `go test` command:
cd greeting_http_server/handlers
go test


## Interacting with the HTTP Server

### Creating a New Item

To create a new message, use the `POST` method with the following `curl` command. Ensure the ID length is greater than 1:
curl "http://localhost:9010/greeting-message" -d '{"id":"1234","name":"test"}' -X POST


### Updating an Existing Item

To update an existing message, use the `PUT` method with the following `curl` command. Replace `1234` with the ID of the item you wish to update:
curl "http://localhost:9010/greeting-message/1234"  -X PUT -d '{"id":"1234","name":"newupdatevalue"}'


### Listing All Items

To list all messages, use the `GET` method with the following `curl` command:

curl "http://localhost:9010/greeting-message" -X GET


### Deleting an Item

To delete a message, use the `DEL` method with the following `curl` command:
curl "http://localhost:9010/greeting-message/1234" -X DEL  

That's it! You're now equipped to interact with the HTTP server in your local development environment. Happy coding!