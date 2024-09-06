# Build stage
FROM golang:1.23.1-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the entire project
COPY . .

# Download dependencies
RUN go mod download

# Build the application
RUN go build -o ipinfo-app

# Final stage
FROM alpine:latest

WORKDIR /root/

# Copy the binary from the builder stage
COPY --from=builder /app/ipinfo-app .

# Copy the internal directory (for templates and other assets)
COPY --from=builder /app/internal ./internal

# Expose the port the app runs on
EXPOSE 8080

# Run the binary
CMD ["./ipinfo-app"]