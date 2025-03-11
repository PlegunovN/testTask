
FROM golang:1.22-alpine AS builder


RUN apk add --no-cache git

WORKDIR /src


COPY go.mod go.sum ./
RUN go mod download


COPY . .


RUN CGO_ENABLED=0 GOOS=linux go build -o grpc_service ./cmd/server


FROM alpine:latest

WORKDIR /app


COPY --from=builder /src/grpc_service .
COPY --from=builder /src/.env .


# 50051 - для gRPC сервиса
# 2112  - для экспонирования метрик
EXPOSE 50051 2112


CMD ["./grpc_service"]

