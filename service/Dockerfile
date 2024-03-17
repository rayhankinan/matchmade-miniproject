FROM golang:1.22 as builder

WORKDIR /server

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o service .

# Final stage
FROM alpine:latest

WORKDIR /root/

# Copy the binary from the builder stage
COPY --from=builder /server/service .

EXPOSE 8080

# Run the application
CMD ["./service", "start"]
