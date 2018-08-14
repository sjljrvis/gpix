package models

import (
	tools "github.com/sjljrvis/gpix/tools"
	"gopkg.in/mgo.v2/bson"
)

const (
	collection = "user"
)

// User model schema
type User struct {
	ID       bson.ObjectId `bson:"_id" json:"_id"`
	UserName string        `bson:"userName" json:"userName"`
}

/*
FindAll runs empty search and returns all documents from collection
*/
func FindAll() ([]User, error) {
	var results []User
	err := tools.MongoDB.C(collection).Find(nil).All(&results)
	return results, err
}
