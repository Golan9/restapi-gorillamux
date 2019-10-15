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

// User struct ..
type User struct {
	UserID   string `bson:"userID" json:"userID"`
	Username string `bson:"username" json:"username"`
	Password string `bson:"password" json:"password"`
	// Token    string `bson:"token" json:"token"`
	Status Status `bson:"status" json:"status"`
}

// LoginResponse struct ..
type LoginResponse struct {
	Username string `bson:"username" json:"username"`
	Password string `bson:"password" json:"password"`
	Token    string `bson:"token" json:"token"`
}

type HttpResponse struct {
	Error   string
	Message string
	Data    string
}

// Users Slice Data ..
// A Slice is basically a variable length array
// user struct as a slice
type Users struct {
	Users []User
}
