package mongodb

import (
    "log"
    "gopkg.in/mgo.v2"
)

const (
    MongoServerAddr = "localhost:27017"
)

var MongoDBSession *mgo.Session

func init() {
    var err error
    MongoDBSession, err = mgo.Dial(MongoServerAddr)
    if err != nil {
		errString := "Error connecting to MongoDB"
		log.Printf(errString)
		panic(errString)
	}
    MongoDBSession.SetSafe(&mgo.Safe{})
}