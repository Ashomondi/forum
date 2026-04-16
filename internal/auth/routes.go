package auth

import "net/http"

func RegisterRoutes(handler *Handler) {
	http.HandleFunc("/register", handler.Register)
	http.HandleFunc("/login", handler.Login)
}
