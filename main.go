package main

import (
	"fmt"
	router "github.com/sjljrvis/gpix/router"
	"github.com/subosito/gotenv"
	"log"
	"net/http"
	"os"
	"text/template"
)

func init() {
	gotenv.Load()
	log.Print("Server started at ", os.Getenv("PORT"))
	fmt.Print("{{1}}")
}

func main() {
	r := router.NewRouter()
	http.Handle("/", r)
	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}
