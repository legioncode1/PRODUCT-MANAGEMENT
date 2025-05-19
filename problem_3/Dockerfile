# Use the official Golang image as a build stage
FROM golang:1.20 AS builder

# Set the current working directory inside the container
WORKDIR /app

# Copy the go.mod and go.sum files
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy the source code into the container
COPY . .

# Set environment variables for architecture
ENV GOARCH=amd64
ENV GOOS=linux

# Build the Go app and list files to confirm the binary is created
RUN go build -o main . && ls -l

# Start a new stage from scratch
FROM alpine:latest  

# Set the current working directory
WORKDIR /root/

# Copy the binary from the builder stage
COPY --from=builder /app/main .

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./main"]

