# syntax=docker/dockerfile:1

# Build stage
FROM golang:1.22 as builder

WORKDIR /app

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o service .

# Final stage
FROM alpine:latest

WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/service .

EXPOSE 8080

# Run the application
CMD ["./service", "start"]
