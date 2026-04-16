package auth

import (
	"encoding/json"
	"net/http"
)

type Handeler struct {
	Service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{Service: service}
}

func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {
	var user User
	json.NewDecoder(r.Body).Decode(&user)
	err := h.Service.Register(user)
	if err != nil {
		http.Error(w, "Cannot register user", http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("User registered successfully"))
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var user User
	json.NewDecoder(r.Body).Decode(&user)

	loggedUser, err := h.Service.login(user.Email, user.Password)
	if err != nill {
		http.Error(w, "Invalid credentials, http.StatusUnauthorized")
		return
	}
	json.NewEncoder(w).Encode(loggedUser)
}
