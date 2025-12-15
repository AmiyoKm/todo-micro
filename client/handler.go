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
