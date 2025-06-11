# Stage 1 — Build the Go binary
FROM golang:1.24-alpine AS builder

# Enable Go modules
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64

WORKDIR /app

# Download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN go build -o main .
# Stage 2 — Run the app with a minimal image
FROM alpine:latest

WORKDIR /app

# Copy binary from builder stage
COPY --from=builder /app/main .


EXPOSE 8080

ENTRYPOINT ["./main"]
