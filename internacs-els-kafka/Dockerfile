# Build stage
FROM golang:1.22.5-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o main ./cmd/api

# Runtime stage
FROM alpine:3.18

# Set environment variable
ENV PORT=8080
ENV GIN_MODE=release

# Copy the pre-built binary from the builder stage
COPY --from=builder /app/main /main

# Expose the port
EXPOSE 8080

# Command to run the executable
CMD ["/main"]