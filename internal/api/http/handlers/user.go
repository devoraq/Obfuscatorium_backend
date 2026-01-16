package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/devoraq/Obfuscatorium_backend/internal/api/http/dto"
	"github.com/devoraq/Obfuscatorium_backend/internal/domain/models"
	"github.com/devoraq/Obfuscatorium_backend/internal/domain/usecases"
)

type UserHandler struct {
	uc *usecases.UserUseCase
}

func NewUserHandler(userUC *usecases.UserUseCase) *UserHandler {
	return &UserHandler{uc: userUC}
}

func (h *UserHandler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	var req dto.RegisterUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	ctx := r.Context()

	newUser := &models.User{
		Username: req.Username,
		Password: req.Password,
		Email:    req.Email,
	}

	_, err := h.uc.CreateUser(ctx, newUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}

func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	var req dto.GetUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	ctx := r.Context()

	user, err := h.uc.GetUser(ctx, req.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := dto.ToUserResponse(user)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
		return
	}
}
