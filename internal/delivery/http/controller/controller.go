package controller

import (
	"encoding/json"
	"golang-project-template/internal/models"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type UserController struct {
	userUsecase models.UserUsecase
}

func NewUserController(userUsecase models.UserUsecase) *UserController {
	return &UserController{
		userUsecase: userUsecase,
	}
}

func (c *UserController) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var newUser models.User
	if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if _, err := c.userUsecase.CreateUser(&newUser); err != nil {
		http.Error(w, "Failed to create user: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newUser)
}

func (c *UserController) GetByIDHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	userID, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Invalid or missing user ID parameter", http.StatusBadRequest)
		return
	}

	user, err := c.userUsecase.GetUserByID(userID)
	if err != nil {
		http.Error(w, "Failed to retrieve user: "+err.Error(), http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

func (c *UserController) GetAllUsersHandler(w http.ResponseWriter, r *http.Request) {
	users, err := c.userUsecase.GetAllUsers()
	if err != nil {
		http.Error(w, "Failed to retrieve all users: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)

}

func (c *UserController) UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	updatedUser, err := c.userUsecase.UpdateUser(&user)
	if err != nil {
		http.Error(w, "Failed to update the user: "+err.Error(), http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updatedUser)
}

func (c *UserController) DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	userID, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Invalid or missing user ID parameter", http.StatusBadRequest)
		return
	}

	err = c.userUsecase.DeleteUser(userID)
	if err != nil {
		http.Error(w, "Failed to delete user: "+err.Error(), http.StatusInternalServerError)
		return
	}
	res := map[string]string{"message": "User has been successfully deleted"}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}
