FROM golang:1.25.5-alpine3.23 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o user ./cmd/api

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/user .

EXPOSE 8081

CMD ["./user"]
