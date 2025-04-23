FROM golang:1.24-alpine AS builder

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download

COPY internal/ ./internal/
COPY pkg/ ./pkg/


RUN CGO_ENABLED=0 GOOS=linux go build -o /text-service ./internal/services/text-service/main.go

# --- ребят, мультистейдж юзайте пжж) ---
FROM alpine:latest

WORKDIR /app

COPY --from=builder /text-service .
COPY config/ ./config/
COPY db/ ./db/


EXPOSE 50052

CMD ["./text-service"]
