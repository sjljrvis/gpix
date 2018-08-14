package main

import (
	. "github.com/sjljrvis/gpix/logger"
	router "github.com/sjljrvis/gpix/router"
	tools "github.com/sjljrvis/gpix/tools"
	"github.com/subosito/gotenv"
	"net/http"
	"os"
)

func init() {
	gotenv.Load()
	Logger.Info("Loaded env file")
	tools.MongoConnect(os.Getenv("DB_URI"), os.Getenv("DB_NAME"))
	Logger.Info("Server started at :", os.Getenv("PORT"))
}

func main() {
	r := router.NewRouter()
	http.Handle("/", r)
	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}
