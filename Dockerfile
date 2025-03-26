# Use the official Golang image as the base image
FROM golang:1.20-alpine as builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go modules and install dependencies
COPY go.mod go.sum ./
RUN go mod tidy

# Copy the source code to the container
COPY . .

# Build the Go application
RUN go build -o book-api .

# Use a smaller image to run the Go app
FROM alpine:latest

# Install required dependencies
RUN apk --no-cache add ca-certificates

# Copy the built binary from the builder image
COPY --from=builder /app/book-api /usr/local/bin/book-api

# Expose the port your app will run on (use the port your API listens to)
EXPOSE 8081

# Define the command to run your app
CMD ["book-api"]
