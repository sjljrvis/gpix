package router

import (
	"github.com/gorilla/mux"
	UserController "github.com/sjljrvis/gpix/controller/user"
)

// NewRouter is router pointer
func NewRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", UserController.HomePage).Methods("GET")
	return r
}
