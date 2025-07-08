# Build stage
FROM golang:1.21-alpine AS builder

# Set working directory
WORKDIR /app

# Copy go mod files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o gid .

# Final stage
FROM alpine:latest

# Create non-root user
RUN addgroup -g 1000 -S appgroup && \
    adduser -u 1000 -S appuser -G appgroup

# Set working directory
WORKDIR /home/appuser

# Copy binary from builder stage
COPY --from=builder /app/gid /usr/local/bin/gid

# Change ownership and make executable
RUN chown appuser:appgroup /usr/local/bin/gid && \
    chmod +x /usr/local/bin/gid

# Switch to non-root user
USER appuser

# Set entrypoint
ENTRYPOINT ["gid"]

# Default command
CMD ["--help"]
