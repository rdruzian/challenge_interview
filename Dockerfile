# Multi-stage build for Go 1.25
FROM golang:1.25 AS builder
WORKDIR /app

# Cache go mod
COPY go.mod go.sum ./
RUN go mod download

# Copy source
COPY . .

# Build
ENV CGO_ENABLED=0
ENV GOOS=linux
RUN go build -o /app/bin/app ./

# Runtime image
FROM alpine:3.19
WORKDIR /app

# CA certificates and curl for healthcheck
RUN apk add --no-cache ca-certificates curl && update-ca-certificates

COPY --from=builder /app/bin/app /app/app

# Expose port
EXPOSE 8080

# Env defaults (can be overridden in compose)
ENV DB_HOST=localhost
ENV DB_PORT=25432
ENV DB_USER=admin
ENV DB_NAME=challenge_db
ENV DB_PASSWORD=123456

# Healthcheck
HEALTHCHECK --interval=10s --timeout=5s --retries=5 CMD curl -fsS http://localhost:8080/health || exit 1

CMD ["/app/app"]