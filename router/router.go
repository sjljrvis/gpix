package router

import (
	"github.com/gorilla/mux"
	UserController "github.com/sjljrvis/gpix/controller/user"
	"net/http"
)

// NewRouter is router pointer
func NewRouter() *mux.Router {
	fs := http.FileServer(http.Dir("./public/"))
	r := mux.NewRouter()

	/*
		user subrouter
		handle  REST-api /user here
	*/

	userRouter := r.PathPrefix("/api/v1/user").Subrouter()
	userRouter.HandleFunc("/", UserController.GetAll).Methods("GET")
	userRouter.HandleFunc("/{id}", UserController.GetOne).Methods("GET")
	userRouter.HandleFunc("/", UserController.Create).Methods("POST")

	r.PathPrefix("/public/").Handler(http.StripPrefix("/public/", fs))
	return r
}
