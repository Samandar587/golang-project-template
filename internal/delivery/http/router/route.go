package router

import (
	"golang-project-template/internal/delivery/http/controller"
	"net/http"

	"github.com/gorilla/mux"
)

func NewRouter(userController *controller.UserController) http.Handler {
	router := mux.NewRouter()

	// Define routes and associate them with handler functions
	router.HandleFunc("/users", userController.CreateUserHandler).Methods("POST")
	router.HandleFunc("/users/{id}", userController.GetByIDHandler).Methods("GET")
	router.HandleFunc("/users", userController.GetAllUsersHandler).Methods("GET")
	router.HandleFunc("/users", userController.UpdateUserHandler).Methods("PUT")
	router.HandleFunc("/users/{id}", userController.DeleteUserHandler).Methods("DELETE")

	return router
}
