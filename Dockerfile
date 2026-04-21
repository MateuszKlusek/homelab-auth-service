# Build Stage
FROM golang:1.24-alpine AS builder
WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o auth-service ./cmd/server/main.go

# Run Stage
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/auth-service .
EXPOSE 8080
CMD ["./auth-service"]