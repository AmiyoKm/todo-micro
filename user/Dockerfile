FROM golang:1.25.5-alpine3.23

WORKDIR /app

# Install make and other necessary tools
RUN apk add --no-cache make curl

# Install golang-migrate
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.17.0/migrate.linux-amd64.tar.gz | tar xvz && \
	mv migrate /usr/local/bin/migrate && \
	chmod +x /usr/local/bin/migrate

COPY . .

RUN go build -o user ./cmd/api

CMD ["./user"]
