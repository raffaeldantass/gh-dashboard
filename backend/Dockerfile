# Development stage
FROM golang:1.22-alpine AS dev

WORKDIR /app

RUN apk add --no-cache wget
RUN wget -O /usr/local/bin/air https://github.com/cosmtrek/air/releases/download/v1.49.0/air_1.49.0_linux_amd64 \
    && chmod +x /usr/local/bin/air

# Copy go mod file
COPY go.mod ./

# Initialize module and download dependencies
RUN go mod tidy

# Copy source code
COPY . .

# Run air for development
CMD ["air"]

# Production stage
FROM golang:1.22-alpine AS builder

WORKDIR /app

# Copy go mod file
COPY go.mod ./

# Initialize module and download dependencies
RUN go mod tidy

# Copy source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/main

# Final production stage
FROM alpine:latest AS prod

WORKDIR /app

# Copy the binary from builder
COPY --from=builder /app/main .

# Run the application
CMD ["./main"]