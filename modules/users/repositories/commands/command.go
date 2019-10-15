package commands

import (
	"log"

	m "restapi-gorillamux/modules/users/models"

	database "restapi-gorillamux/database"
)

const (
	// COLLECTION ..
	COLLECTION = "users"
)

// Commands struct ..
type Commands struct{}

// Insert ..
func (c *Commands) Insert(user m.User) error {
	//fmt.Println("masuk insert", user)
	err := database.Db.C(COLLECTION).Insert(&user)
	//fmt.Println("cek error", err)

	if err != nil {
		log.Println(err, "Error Insert")
	}
	return err
}
