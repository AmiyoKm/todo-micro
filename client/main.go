package main

import (
	"fmt"
	"net/http"
	"os"
)

var rabbitMQ *RabbitMQ

func main() {
	url := os.Getenv("RABBITMQ_URL")
	rabbit, err := NewRabbitMQ(url)
	if err != nil {
		fmt.Println("Error connecting to RabbitMQ:", err)
		os.Exit(1)
	}
	defer rabbitMQ.Close()

	rabbitMQ = rabbit

	mux := http.NewServeMux()
	mux.HandleFunc("GET /todos/{id}", getTodo)
	mux.HandleFunc("POST /todos", createTodo)
	mux.HandleFunc("PATCH /todos/{id}", updateTodo)
	mux.HandleFunc("DELETE /todos/{id}", deleteTodo)
	http.ListenAndServe(":3000", mux)
}
