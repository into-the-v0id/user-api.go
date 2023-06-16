package route

import (
	"github.com/gorilla/mux"
	"github.com/into-the-v0id/user-api.go/controller"
)

func RegisterUserRoutes(router *mux.Router) {
	router.HandleFunc("/", controller.UserList).Methods("GET")
	router.HandleFunc("/", controller.UserCreate).Methods("POST")
	router.HandleFunc("/{id}", controller.UserDetail).Methods("GET")
	router.HandleFunc("/{id}", controller.UserUpdate).Methods("PUT")
	router.HandleFunc("/{id}", controller.UserDelete).Methods("DELETE")
}
