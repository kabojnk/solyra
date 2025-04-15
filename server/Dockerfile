FROM golang:1.24-alpine AS builder

# Install git and air for hot reloading
RUN apk add --no-cache git && go install github.com/air-verse/air@latest

WORKDIR /app

# Copy go mod and sum files
COPY cmd/api/go.mod cmd/api/go.sum ./cmd/api/

# Download dependencies
RUN go mod download

# Copy source code
COPY .. .

# Build the application
RUN go build -o /app/etrenank/api ./cmd/api/main.go

FROM golang:1.24-alpine

# Copy the binary from builder
COPY --from=builder /app/etrenank/api .
COPY --from=builder /go/bin/air /usr/local/bin/air

# Copy air configuration
COPY .air.toml .

# Expose port
EXPOSE 8080

# Set environment variables
ENV GIN_MODE=release

# Run with air for hot reloading in development
CMD ["air", "-c", ".air.toml"]
