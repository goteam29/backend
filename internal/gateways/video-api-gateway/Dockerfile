FROM golang:1.24-alpine AS builder

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod tidy
RUN go mod download

COPY internal/ ./internal/
COPY pkg/ ./pkg/


RUN CGO_ENABLED=0 GOOS=linux go build -o /video-gateway ./internal/gateways/video-api-gateway/main.go


FROM alpine:latest

WORKDIR /app

COPY --from=builder /video-gateway .
COPY config/ ./config/

EXPOSE 8081

CMD ["./video-gateway"]
