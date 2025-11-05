# Build stage
FROM golang:1.25.3-alpine AS builder

WORKDIR /app

# Install dependencies
RUN apk add --no-cache git

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o /school-management-system ./cmd/server

# Final stage
FROM alpine:latest

WORKDIR /app

# Install necessary runtime dependencies
RUN apk --no-cache add ca-certificates

# Copy the binary from builder
COPY --from=builder /school-management-system .

# Copy environment file (if exists)
COPY .env* ./

# Expose the application port
EXPOSE 8080

# Command to run the application
CMD ["./school-management-system"]
