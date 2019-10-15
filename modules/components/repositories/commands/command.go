package commands

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

// Commands struct ..
type Commands struct{}

// Test ..
func Test() {
	fmt.Println("TEST IMPORT")
}

// Insert ..
func (c *Commands) Insert(component m.Component) error {
	err := database.Db.C(COLLECTION).Insert(&component)
	if err != nil {
		log.Println(err, "Error Insert")
	}
	return err
}

// Delete ..
func (c *Commands) Delete(componentID string, component m.Component) error {
	err := database.Db.C(COLLECTION).Update(bson.M{"componentID": componentID}, bson.M{"$set": bson.M{"component_name": component.ComponentName,
		"component_desc": component.ComponentDesc,
		"status":         component.Status,
		"data":           component.ComponentList}})

	if err != nil {
		log.Println(err, "Error Delete")
	}
	return err
}

// Update ..
func (c *Commands) Update(componentID string, component m.Component) error {
	// var component m.Component
	err := database.Db.C(COLLECTION).Update(bson.M{"componentID": componentID}, bson.M{"$set": bson.M{"component_name": component.ComponentName,
		"component_desc": component.ComponentDesc,
		"status":         component.Status,
		"data":           component.ComponentList}})

	if err != nil {
		log.Println(err, "Error Update")
	}
	return err
}
