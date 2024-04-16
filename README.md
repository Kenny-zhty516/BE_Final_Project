# Welcome to your CDK Python project!

## Local dev testing

Use Docker Compose to start a local development and testing environment.

```bash
cd local_dev
make
```

## List tables in local DynamoDB

```bash
aws dynamodb list-tables --endpoint-url http://localhost:8000
```

## List items in a given table in local DynamoDB

```bash
aws dynamodb scan --table-name GreetingMessage --endpoint-url http://localhost:8000
```

## Invoke Lambda locally

```bash
curl "http://localhost:9000/2015-03-31/functions/function/invocations" -d '{"id": "1234", "name": "Amos"}'
```