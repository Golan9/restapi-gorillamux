package queries

import (
	"fmt"
	"log"

	database "restapi-gorillamux/database"
	m "restapi-gorillamux/modules/components/models"

	"gopkg.in/mgo.v2/bson"
)

const (
	// COLLECTION ..
	COLLECTION = "components"
)

// Queries struct ..
type Queries struct{}

// Test ..
func Test() {
	fmt.Println("TEST IMPORT")
}

// FindAll ..
func (q *Queries) FindAll() ([]m.Component, error) {
	components := m.Components{}
	err := database.Db.C(COLLECTION).Find(bson.M{}).All(&components.Components)
	if err != nil {
		log.Println(err, "Error Find All")
	}
	return components.Components, err
}

// FindByID ..
func (q *Queries) FindByID(componentID string) (m.Component, error) {
	var component m.Component
	err := database.Db.C(COLLECTION).Find(bson.M{"componentID": componentID}).One(&component)
	if err != nil {
		log.Println(err, "Invalid ComponentID!")
	}
	return component, err
}
