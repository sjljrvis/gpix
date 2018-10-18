package controller

import (
	"encoding/json"
	"github.com/gorilla/mux"
	Helper "github.com/sjljrvis/gpix/helper"
	BillModel "github.com/sjljrvis/gpix/models/bill"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"time"
)

// GetAll controller
func GetAll(w http.ResponseWriter, r *http.Request) {
	var results []BillModel.Bill
	results, err := BillModel.FindAll()
	if err != nil {
		Helper.RespondWithError(w, 200, err.Error()) // TODO: Do something about the error
	} else {
		Helper.RespondWithJSON(w, 200, results)
	}
}

// GetOne controller
func GetOne(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	result, err := BillModel.FindOneByID(params["id"])
	if err != nil {
		Helper.RespondWithError(w, 200, err.Error()) // TODO: Do something about the error
	} else {
		Helper.RespondWithJSON(w, 200, result)
	}
}

// Create controller
func Create(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var bill BillModel.Bill

	if err := json.NewDecoder(r.Body).Decode(&bill); err != nil {
		Helper.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	bill.ID = bson.NewObjectId()
	bill.GeneratedDate = time.Now()
	err := BillModel.Create(bill)
	if err != nil {
		Helper.RespondWithError(w, 200, err.Error()) // TODO: Do something about the error
	} else {
		Helper.RespondWithJSON(w, 200, map[string]string{"message": "Bill Created successfully"})
	}
}

// Search controller
func Search(w http.ResponseWriter, r *http.Request) {

	var query map[string]interface{}
	query = make(map[string]interface{})

	keys := r.URL.Query()

	for item := range keys {
		query[item] = keys[item][0]
	}

	result, err := BillModel.FindByQuery(query)
	if err != nil {
		Helper.RespondWithError(w, 200, err.Error()) // TODO: Do something about the error
	} else {
		Helper.RespondWithJSON(w, 200, result)
	}
}
