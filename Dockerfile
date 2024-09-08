ARG GO_VERSION=1.22
FROM golang:${GO_VERSION}-alpine as builder

WORKDIR /app

ADD go.mod ./
RUN go mod download

ADD . ./
RUN go build -o /go-app

# Second stage
FROM alpine:latest
WORKDIR /root/

EXPOSE 8080

CMD ["./go-app"]
