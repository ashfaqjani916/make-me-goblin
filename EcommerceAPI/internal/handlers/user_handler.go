package handlers

import (
	"EcommerceAPI/internal/models"
	"EcommerceAPI/internal/services"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type UserHandler struct{
	service services.UserService
}

func NewUserHandler(service services.UserService)  *UserHandler {
	return &UserHandler{service}
}


func (h *UserHandler) GetAllUsers(res http.ResponseWriter, req *http.Request) {
	users, err := h.service.GetAllUsers()
	if err!= nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(res).Encode(users)
}


func (h *UserHandler) GetUserByID(res http.ResponseWriter, req *http.Request) {
	param := mux.Vars(req)
	
	id, err := strconv.Atoi(param["id"])
	if err != nil {
		http.Error(res, "Invalid ID", http.StatusBadRequest)
		return
	}

user, err := h.service.GetUserById(uint(id))
	if err != nil {
		http.Error(res, "User not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(res).Encode(user)
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	err := h.service.CreateUser(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}
