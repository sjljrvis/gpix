package router

import (
	"github.com/gorilla/mux"

	BillController "github.com/sjljrvis/gpix/controller/bill"
	OauthController "github.com/sjljrvis/gpix/controller/oauth"
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
	userRouter.HandleFunc("/search/", UserController.Search).Methods("GET")

	/*
		bill subrouter
		handle  REST-api /user here
	*/

	billsRouter := r.PathPrefix("/api/v1/bill").Subrouter()
	billsRouter.HandleFunc("/", BillController.GetAll).Methods("GET")
	billsRouter.HandleFunc("/{id}", BillController.GetOne).Methods("GET")
	billsRouter.HandleFunc("/", BillController.Create).Methods("POST")
	billsRouter.HandleFunc("/search/", BillController.Search).Methods("GET")

	oauthRouter := r.PathPrefix("/api/v1/oauth").Subrouter()
	oauthRouter.HandleFunc("/", OauthController.OauthHandler).Methods("GET")
	oauthRouter.HandleFunc("/google/login", OauthController.GloginHandler).Methods("GET")
	oauthRouter.HandleFunc("/google/callback", OauthController.GcallbackHandler).Methods("GET")

	r.PathPrefix("/public/").Handler(http.StripPrefix("/public/", fs))
	return r
}
