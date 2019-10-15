package models

// To store struct
// Struct is like a class, it is used for oop in golang
// (properties, methods) silimar tu ES6, Java, C# etc.

// Status ..
type Status int32

const (

	// Active is status when component is active
	Active Status = 1

	// InActive is status when component is inactive
	InActive Status = -1

	// Deleted is status when component is deleted
	Deleted Status = -2
)

// Component Struct ..
type Component struct {
	ComponentID   string `bson:"componentID" json:"componentID"`
	ComponentName string `bson:"component_name" json:"component_name"`
	ComponentDesc string `bson:"component_desc" json:"component_desc"`
	Status        Status `bson:"status" json:"status"`
	ComponentList *Data  `bson:"data" json:"data"`
}

// Data Struct ..
type Data struct {
	Usage   string `bson:"usage" json:"usage"`
	Anatomy string `bson:"anatomy" json:"anatomy"`
	Specs   string `bson:"specs" json:"specs"`
}

// Components Slice Data ..
// A Slice is basically a variable length array
// component struct as a slice
type Components struct {
	Components []Component
}
