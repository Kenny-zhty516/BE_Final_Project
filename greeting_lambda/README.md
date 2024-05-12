# Welcome to your CDK Python project!

## Local dev testing

Use Docker Compose to start a local development and testing environment.

```bash
cd local_dev
make
```

## Invoke Lambda locally

```bash
curl "http://localhost:9000/2015-03-31/functions/function/invocations" -d '{"body":"{\"id\":\"123\",\"name\":\"test\"}"}'
```

Should see response

```
{"statusCode":200,"headers":null,"multiValueHeaders":null,"body":"Item with ID 123 added."}
```

Verify by doing a scan on the table.

## Create new item using HTTP server

To create a message using the HTTP server we use the following command.

```bash
curl "http://localhost:9010/greeting-message" -d '{"id":"1234","name":"test"} -X POST'
```

