package commands

import (
	"log"

	mongo "gopkg.in/mgo.v2"
)

// DAO ..
type DAO struct {
	Server   string
	Database string
}

// Db ..
var Db *mongo.Database

// Connect ..
func (u *DAO) Connect() {
	session, err := mongo.Dial(u.Server)
	if err != nil {
		log.Println(err, "Error connect to mongodb")
	}
	Db = session.DB(u.Database)
}
