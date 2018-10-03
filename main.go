package main

import (
	"github.com/gorilla/handlers"
	. "github.com/sjljrvis/gpix/logger"
	router "github.com/sjljrvis/gpix/router"
	socket "github.com/sjljrvis/gpix/sockets"
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
	hub := socket.NewHub()
	go hub.Run()

	corsObj := handlers.AllowedOrigins([]string{"*"})

	http.Handle("/", handlers.CORS(corsObj)(r))
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		socket.ServeWs(hub, w, r)
	})
	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}
