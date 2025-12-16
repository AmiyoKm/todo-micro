package main

import (
	"database/sql"
	"fmt"
	"os"
)

type app struct {
	rabbitMQ *RabbitMQ
	db       *sql.DB
}

func main() {
	url := os.Getenv("RABBITMQ_URL")
	rabbitMQ, err := NewRabbitMQ(url)
	if err != nil {
		fmt.Println("Failed to connect to RabbitMQ:", err)
		return
	}
	defer rabbitMQ.Close()
	err = rabbitMQ.QueueDeclare("todo")
	if err != nil {
		fmt.Println("Failed to declare queue:", err)
		return
	}
	msgs, err := rabbitMQ.Consume("todo")
	if err != nil {
		fmt.Println("Failed to consume queue:", err)
		return
	}

	dbConfig := DbConfig{
		Host:        os.Getenv("POSTGRES_HOST"),
		User:        os.Getenv("POSTGRES_USER"),
		Password:    os.Getenv("POSTGRES_PASSWORD"),
		Name:        os.Getenv("POSTGRES_DB"),
		Port:        os.Getenv("POSTGRES_PORT"),
		SslMode:     os.Getenv("POSTGRES_SSLMODE"),
		MaxConnOpen: 60,
		MaxIdleConn: 60,
		MaxIdleTime: "10m",
	}
	db, err := NewDb(dbConfig)
	if err != nil {
		fmt.Println("Failed to connect to database:", err)
		return
	}
	defer db.Close()

	app := &app{
		rabbitMQ: rabbitMQ,
		db:       db,
	}

	forever := make(chan struct{})
	go func() {
		for msg := range msgs {
			err := app.handleMessage(msg)
			if err != nil {
				fmt.Println("Failed to handle message:", err)
			}
		}
	}()

	fmt.Println("Started consuming messages")

	<-forever
}
