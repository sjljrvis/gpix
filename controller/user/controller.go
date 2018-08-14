package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	Helper "github.com/sjljrvis/gpix/helper"
	UserModel "github.com/sjljrvis/gpix/models/user"
	"gopkg.in/mgo.v2/bson"
	"net/http"
)

// GetAll controller
func GetAll(w http.ResponseWriter, r *http.Request) {
	var results []UserModel.User
	results, err := UserModel.FindAll()
	if err != nil {
		Helper.RespondWithError(w, 200, err.Error()) // TODO: Do something about the error
	} else {
		Helper.RespondWithJSON(w, 200, results)
	}
}

// GetOne controller
func GetOne(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	result, err := UserModel.FindOneByID(params["id"])
	if err != nil {
		Helper.RespondWithError(w, 200, err.Error()) // TODO: Do something about the error
	} else {
		Helper.RespondWithJSON(w, 200, result)
	}
}

// Create controller
func Create(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var user UserModel.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		Helper.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	user.ID = bson.NewObjectId()
	err := UserModel.Create(user)
	if err != nil {
		Helper.RespondWithError(w, 200, err.Error()) // TODO: Do something about the error
	} else {
		Helper.RespondWithJSON(w, 200, map[string]string{"message": "User Created successfully"})
	}
}

// SearchUser controller
func SearchUser() {
	_err := UserModel.FindByQuery(map[string]interface{}{"userName": "jarvis"})
	fmt.Println("---------", _err, "--------------")
}
