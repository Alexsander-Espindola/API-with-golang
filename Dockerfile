FROM golang:1.18 AS build

WORKDIR /src

COPY go.mod ./
COPY src/aprendendo.go ./

RUN go build -o /server

EXPOSE 8080

ENTRYPOINT ["/server"]
