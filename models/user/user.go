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
	ID       bson.ObjectId          `bson:"_id" json:"_id"`
	UserName string                 `bson:"userName" json:"userName"`
	Email    string                 `bson:"email" json:"email"`
	Shopify  map[string]interface{} `bson:"shopify" json:"shopify"`
	PlanType string                 `bson:"planType" json:"planType"`
}

/*
FindAll runs empty search and returns all documents from collection
*/
func FindAll() ([]User, error) {
	var results []User
	err := tools.MongoDB.C(collection).Find(nil).All(&results)
	return results, err
}

/*
FindOneByID runs query based on _id  and returns single document from collection
*/
func FindOneByID(id string) (User, error) {
	var result User
	err := tools.MongoDB.C(collection).FindId(bson.ObjectIdHex(id)).One(&result)
	return result, err
}

/*
Create runs query based on _id  and returns single document from collection
*/
func Create(user User) error {
	err := tools.MongoDB.C(collection).Insert(&user)
	return err
}

/*
FindByQuery runs query based on query  and returns array of document from collection
*/
func FindByQuery(query map[string]interface{}) ([]User, error) {
	var results []User
	err := tools.MongoDB.C(collection).Find(query).All(&results)
	return results, err
}
