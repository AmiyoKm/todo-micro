package main

import (
	"encoding/json"
	"net/http"
)

func getTodo(w http.ResponseWriter, r *http.Request) {

	id := r.PathValue("id")

	todoService, err := NewTodoServerClient()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer todoService.Close()

	payload := GetTodoRequest{
		Id: id,
	}

	todo, err := todoService.Client.GetTodo(r.Context(), payload.toProto())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(todo); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func createTodo(w http.ResponseWriter, r *http.Request) {
	var payload CreateTodoRequest

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	todoService, err := NewTodoServerClient()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer todoService.Close()

	todo, err := todoService.Client.CreateTodo(r.Context(), payload.toProto())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(todo); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

type UpdateTodoRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
}

func updateTodo(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	ctx := r.Context()
	var payload UpdateTodoRequest

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	message := &Message{
		Topic: "todo.update",
		Todo: &Todo{
			Id:          id,
			Title:       payload.Title,
			Description: payload.Description,
			Done:        payload.Done,
		},
	}

	if err := rabbitMQ.Publish(ctx, message); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(message); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func deleteTodo(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	ctx := r.Context()

	message := &Message{
		Topic: "todo.delete",
		Todo: &Todo{
			Id: id,
		},
	}

	if err := rabbitMQ.Publish(ctx, message); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(message); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
