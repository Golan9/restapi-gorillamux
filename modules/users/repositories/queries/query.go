package queries

import (
	"fmt"
	"log"

	database "restapi-gorillamux/database"
	m "restapi-gorillamux/modules/users/models"

	"gopkg.in/mgo.v2/bson"
)

// Queries struct ..
type Queries struct{}

const (
	// COLLECTION ..
	COLLECTION = "users"
)

// Test ..
func Test() {
	fmt.Println("TEST IMPORT")
}

// FindAll ..
func (q *Queries) FindAll() ([]m.User, error) {
	users := m.Users{}
	err := database.Db.C(COLLECTION).Find(bson.M{}).All(&users.Users)
	if err != nil {
		log.Println(err, "Error Find All")
	}
	return users.Users, err
}

// FindByID ..
func (q *Queries) FindByID(userID string) (m.User, error) {
	var user m.User
	err := database.Db.C(COLLECTION).Find(bson.M{"userID": userID}).One(&user)
	if err != nil {
		log.Println(err, "Invalid UserID!")
	}
	return user, err
}

// FindOne ..
func (q *Queries) FindOne(username string) (m.User, error) {
	var user m.User
	err := database.Db.C(COLLECTION).Find(bson.M{"username": username}).One(&user)
	if err != nil {
		log.Println(err, "Invalid Username!")
	}
	return user, err
}
