package main

import (
	router "github.com/sjljrvis/gpix/router"
	"log"
	"net/http"
)

func init() {
	log.Print("Server started")
}

func main() {
	r := router.NewRouter()
	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}
