
# Use an official Go image as the builder stage
FROM golang:1.23.5-alpine AS builder

WORKDIR /app

# Copy go.mod and go.sum and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the entire source code
COPY . .

# Build the Go binary
RUN go build -o TodoAPI .

# Use a minimal base image for the final stage
FROM alpine:latest

WORKDIR /root/

# Copy the compiled binary from the builder stage
COPY --from=builder /app/TodoAPI .

# Expose the application port
EXPOSE 8080

# Run the application
CMD ["./TodoAPI"]
