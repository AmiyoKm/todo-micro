package main

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/AmiyoKm/todo-micro/api-gateway/gen/userpb"
)

func createUser(w http.ResponseWriter, r *http.Request) {
	var payload CreateUserRequest

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if payload.Email == "" || payload.Name == "" || payload.Password == "" {
		http.Error(w, "Email, name, and password are required", http.StatusBadRequest)
		return
	}

	ctx := r.Context()

	client, err := NewUserServerClient()
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer client.Close()

	req := &userpb.CreateUserRequest{
		Email:    payload.Email,
		Name:     payload.Name,
		Password: payload.Password,
	}

	user, err := client.Client.CreateUser(ctx, req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

func getUser(w http.ResponseWriter, r *http.Request) {
	header := r.Header.Get("Authorization")
	if header == "" {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	parts := strings.Split(header, " ")
	if len(parts) != 2 || parts[1] == "" {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}
	token := parts[1]

	ctx := r.Context()

	client, err := NewUserServerClient()
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer client.Close()

	req := &userpb.GetUserRequest{
		Jwt: token,
	}

	user, err := client.Client.GetUser(ctx, req)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func login(w http.ResponseWriter, r *http.Request) {
	var payload LoginRequest

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if payload.Email == "" || payload.Password == "" {
		http.Error(w, "Email and password are required", http.StatusBadRequest)
		return
	}

	ctx := r.Context()

	client, err := NewUserServerClient()
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer client.Close()

	req := &userpb.LoginRequest{
		Email:    payload.Email,
		Password: payload.Password,
	}

	loginResp, err := client.Client.Login(ctx, req)
	if err != nil {
		http.Error(w, "Invalid email or password", http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(LoginResponse{
		Jwt: loginResp.Jwt,
	})
}
