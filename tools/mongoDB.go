package tools

import (
	mgo "gopkg.in/mgo.v2"
	"log"
)

//MongoDB instance
var MongoDB *mgo.Database

//MongoConnect -Connecting to DB
func MongoConnect(mongoURI, dbName string) {
	log.Print("Connecting to mongo")
	session, err := mgo.Dial(mongoURI)
	if err != nil {
		log.Fatal(err)
	}
	MongoDB = session.DB(dbName)
}
