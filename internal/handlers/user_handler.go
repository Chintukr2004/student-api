package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Chintukr2004/student-api/internal/service"
)

type UserHandler struct {
	Service   *service.UserService
	JWTSecret string
}

func NewUserHandler(service *service.UserService, jwtSecret string) *UserHandler {
	return &UserHandler{Service: service,
		JWTSecret: jwtSecret}
}

type loginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {

	var req loginRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Println("DEBUG HANDLER: JSON decode failed", err)
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	token, err := h.Service.Login(
		r.Context(),
		req.Email,
		req.Password,
		h.JWTSecret,
	)

	if err != nil {
		http.Error(w, "invalid credentials", http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(map[string]string{
		"token": token,
	})
}

type registerRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	var req registerRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	user, err := h.Service.Register(
		r.Context(),
		req.Name,
		req.Email,
		req.Password,
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}
