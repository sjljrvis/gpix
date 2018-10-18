package models

import (
	tools "github.com/sjljrvis/gpix/tools"
	"gopkg.in/mgo.v2/bson"
	"time"
)

const (
	collection = "bill"
)

// Bill model schema
type Bill struct {
	ID            bson.ObjectId `bson:"_id" json:"_id"`
	UserName      string        `bson:"userName" json:"userName"`
	Email         string        `bson:"email" json:"email"`
	Amount        int64         `bson:"amount" json:"amount"`
	GeneratedDate time.Time     `bson:"generatedDate" json:"generatedDate"`
	IsPaid        bool          `bson:"isPaid" json:"isPaid"`
}

/*
FindAll runs empty search and returns all documents from collection
*/
func FindAll() ([]Bill, error) {
	var results []Bill
	err := tools.MongoDB.C(collection).Find(nil).All(&results)
	return results, err
}

/*
FindOneByID runs query based on _id  and returns single document from collection
*/
func FindOneByID(id string) (Bill, error) {
	var result Bill
	err := tools.MongoDB.C(collection).FindId(bson.ObjectIdHex(id)).One(&result)
	return result, err
}

/*
Create runs query based on _id  and returns single document from collection
*/
func Create(user Bill) error {
	err := tools.MongoDB.C(collection).Insert(&user)
	return err
}

/*
FindByQuery runs query based on query  and returns array of document from collection
*/
func FindByQuery(query map[string]interface{}) ([]Bill, error) {
	var results []Bill
	err := tools.MongoDB.C(collection).Find(query).All(&results)
	return results, err
}
