FROM golang:1.18 AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY *.go ./

RUN go build -o /docker-golang-grpc
RUN go run src/proto/cmd/server/server.go
RUN go run src/proto/cmd/client/client.go

EXPOSE 8080

CMD [ "/docker-golang-grpc" ]
