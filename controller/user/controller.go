package controller

import (
	helper "github.com/sjljrvis/gpix/helper"
	model "github.com/sjljrvis/gpix/models/user"
	tools "github.com/sjljrvis/gpix/tools"
	"net/http"
)

// HomePage controller
func HomePage(w http.ResponseWriter, r *http.Request) {
	var results []model.User
	err := tools.MongoDB.C("user").Find(nil).All(&results)
	if err != nil {
		helper.RespondWithError(w, 200, "error in mongoDB ") // TODO: Do something about the error
	} else {
		helper.RespondWithJSON(w, 200, results)
	}
}
