FROM golang:1.24-alpine AS builder

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download

COPY internal/ ./internal/
COPY pkg/ ./pkg/
COPY db/migrations/ ./db/migrations/


RUN CGO_ENABLED=0 GOOS=linux go build -o /user-service ./internal/services/user-service/main.go

# --- ребят, мультистейдж юзайте пжж) ---
FROM alpine:latest

WORKDIR /app

COPY --from=builder /user-service .
COPY config/ ./config/
COPY db/migrations/ ./db/migrations/

EXPOSE 50053

CMD ["./user-service"]
