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
	api := r.PathPrefix("/api/v1").Subrouter()

	api.HandleFunc("/", UserController.HomePage).Methods("GET")
	api.PathPrefix("/public/").Handler(http.StripPrefix("/public/", fs))
	return r
}
