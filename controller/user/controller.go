package controller

import (
	helper "github.com/sjljrvis/gpix/helper"
	UserModel "github.com/sjljrvis/gpix/models/user"
	tools "github.com/sjljrvis/gpix/tools"
	"log"
	"net/http"
)

// HomePage controller
func HomePage(w http.ResponseWriter, r *http.Request) {
	log.Print(tools.MongoDB)
	var results []UserModel.User
	err := tools.MongoDB.C("user").Find(nil).All(&results)
	if err != nil {
		helper.RespondWithError(w, 200, "error in mongoDB ") // TODO: Do something about the error
	} else {
		helper.RespondWithJSON(w, 200, results)
	}
}
