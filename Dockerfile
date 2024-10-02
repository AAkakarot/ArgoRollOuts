# Build Stage
FROM golang:1.22 AS builder

# Set the working directory
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code to the container
COPY . .

# Build the Go binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main ./app

# Run Stage
FROM alpine:latest

# Set the working directory
WORKDIR /root/

# Copy the Go binary from the builder stage
COPY --from=builder /app/main .

# Ensure the binary is executable
RUN chmod +x ./main

# Expose port 8080
EXPOSE 8080

# Command to run the binary
CMD ["./main"]
