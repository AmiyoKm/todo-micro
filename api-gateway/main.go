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

	// Public routes (no authentication required)
	mux.HandleFunc("POST /register", createUser)
	mux.HandleFunc("POST /login", login)

	// Protected routes (authentication required)
	mux.HandleFunc("GET /todos", authMiddleware(getTodos))
	mux.HandleFunc("GET /todos/{id}", authMiddleware(getTodo))
	mux.HandleFunc("POST /todos", authMiddleware(createTodo))
	mux.HandleFunc("PATCH /todos/{id}", authMiddleware(updateTodo))
	mux.HandleFunc("DELETE /todos/{id}", authMiddleware(deleteTodo))
	mux.HandleFunc("GET /users/me", authMiddleware(getUser))

	http.ListenAndServe(":3000", corsMiddleware(http.Handler(mux)))
}
