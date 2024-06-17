package router

import (
	"user-registration/internal/handler"
	"user-registration/internal/service"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter()

	userService := service.NewUserService()
	userHandler := handler.NewUserHandler(userService)

	r.HandleFunc("/register", userHandler.RegisterUser).Methods("POST")

	return r
}
