FROM golang:alpine as build
WORKDIR /usr/src/app
COPY go.mod http-servers.go ./
RUN go build -o /usr/local/bin/http_servers .


FROM alpine as running-app
COPY --from=build /usr/local/bin/http_servers /usr/local/bin/http_servers
EXPOSE 8080
CMD ["http_servers"]