FROM golang:1.25.5-alpine3.23 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o todo ./cmd/api

FROM alpine:latest

WORKDIR /app
COPY --from=builder /app/todo .

EXPOSE 8080

CMD ["./todo"]
