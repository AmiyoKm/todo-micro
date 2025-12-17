FROM golang:1.25.5-alpine3.23

WORKDIR /app

COPY . .

RUN go build -o api-gateway .

CMD ["./api-gateway"]
