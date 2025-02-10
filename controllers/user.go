package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/janjos/user-api/useCases"
)

type UserController struct {
	userUsecase *useCases.UserUsecase
}

func NewUserController(userUsecase *useCases.UserUsecase) *UserController {
	return &UserController{userUsecase: userUsecase}
}

func (c *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Error on requisition decodification", http.StatusBadRequest)
		return
	}

	user, err := c.userUsecase.CreateUser(input.Email, input.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

func (c *UserController) GetUserByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	user, err := c.userUsecase.GetUserByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

func (c *UserController) LogIn(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Error on requisition decodification", http.StatusBadRequest)
		return
	}

	user, err := c.userUsecase.LogIn(input.Email, input.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

func (c *UserController) VerifyToken(w http.ResponseWriter, r *http.Request) {
	type response struct {
		Id float64 `json:"id"`
	}
	token := r.Header.Get("Authorization")
	if token == "" {
		http.Error(w, "No token found", http.StatusUnauthorized)
		return
	}

	id, err := c.userUsecase.VerifyToken(token)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response{Id: id})
}
