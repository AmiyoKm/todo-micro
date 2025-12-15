package main

import "net/http"

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /todos/{id}", getTodo)
	mux.HandleFunc("POST /todos", createTodo)
	http.ListenAndServe(":3000", mux)
}
