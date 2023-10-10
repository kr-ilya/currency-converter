# Start from golang base image
FROM golang:alpine as builder

# Install git. Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git

# Set the current working directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and the go.sum files are not changed 
RUN go mod download

COPY . .

# Build the Go app
RUN go build ./cmd/main.go

# Start a new stage from scratch
FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /opt/bot

# Copy the Pre-built binary file from the previous stage + .env file
COPY --from=builder /app/main .


#Command to run the executable
CMD ["./main"]