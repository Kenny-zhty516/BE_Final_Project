FROM golang:alpine as build
WORKDIR /usr/src/app

# Copy the go.mod and go.sum files first to leverage Docker cache
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the application source code
COPY . .

RUN go build -o /usr/local/bin/http_servers .

FROM alpine as running-app

# Copy the built executable from the builder stage
COPY --from=build /usr/local/bin/http_servers /usr/local/bin/http_servers

# Expose port 8080
EXPOSE 8080

# Command to run
CMD ["http_servers"]