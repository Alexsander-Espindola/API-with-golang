FROM golang:1.18 AS build

ADD . /minipix
WORKDIR /minipix

COPY go.mod ./minipix
COPY ./ ./minipix

RUN go build -o /minipix

EXPOSE 8080

ENTRYPOINT ["/minipix"]
