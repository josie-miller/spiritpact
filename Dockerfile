# First Stage: Build the Go application
ARG GO_VERSION=1.23
FROM golang:${GO_VERSION}-alpine AS builder

WORKDIR /app

# Copy go.mod and go.sum and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the code and build the application
COPY . ./
RUN go build -v -o /go-app

# Second Stage: Final image
FROM alpine:latest

WORKDIR /root/
COPY --from=builder /go-app /go-app

EXPOSE 8080

CMD ["./go-app"]
