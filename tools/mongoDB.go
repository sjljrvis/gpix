package tools

import (
	. "github.com/sjljrvis/gpix/logger"
	mgo "gopkg.in/mgo.v2"
)

//MongoDB instance
var MongoDB *mgo.Database

//MongoConnect -Connecting to DB
func MongoConnect(mongoURI, dbName string) {
	Logger.Info("Connecting to mongo")
	session, err := mgo.Dial(mongoURI)
	if err != nil {
		Logger.Fatal(err)
	}
	MongoDB = session.DB(dbName)
	Logger.Info("Connected to mongo :", dbName)
}
