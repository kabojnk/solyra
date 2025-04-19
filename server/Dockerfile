FROM golang:1.24-alpine AS builder

WORKDIR /build

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o api ./cmd/api/main.go

FROM golang:1.24-alpine

WORKDIR /app

# Copy the binary and configs
COPY --from=builder /build/api .
COPY .env /build/api

# Expose port
EXPOSE 8081

# Set environment variables
ENV GIN_MODE=release

# Run with air for hot reloading in development
CMD ["./api"]
