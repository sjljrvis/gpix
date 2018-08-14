package controller

import (
	helper "github.com/sjljrvis/gpix/helper"
	UserModel "github.com/sjljrvis/gpix/models/user"
	"net/http"
)

// HomePage controller
func HomePage(w http.ResponseWriter, r *http.Request) {
	var results []UserModel.User
	results, err := UserModel.FindAll()
	if err != nil {
		helper.RespondWithError(w, 200, "error in mongoDB ") // TODO: Do something about the error
	} else {
		helper.RespondWithJSON(w, 200, results)
	}
}
