package models

import (
	"gopkg.in/mgo.v2/bson"
)

// User model schema
type User struct {
	ID       bson.ObjectId `bson:"_id" json:"_id"`
	UserName string        `bson:"userName" json:"userName"`
}
