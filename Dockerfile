ARG GO_VERSION=1.22
FROM golang:${GO_VERSION}-alpine as builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./
RUN go build -o /go-app

# Second stage
FROM alpine:latest
WORKDIR /root/

COPY --from=builder /go-app .

EXPOSE 8080

CMD ["./go-app"]
