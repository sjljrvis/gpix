package main

import (
	router "github.com/sjljrvis/gpix/router"
	tools "github.com/sjljrvis/gpix/tools"
	"github.com/subosito/gotenv"
	"log"
	"net/http"
	"os"
)

func init() {
	gotenv.Load()
	log.Print("Loaded env file")
	tools.MongoConnect(os.Getenv("DB_URI"), os.Getenv("DB_NAME"))
	log.Print("Server started at ", os.Getenv("PORT"))
}

func main() {
	r := router.NewRouter()
	http.Handle("/", r)
	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}
