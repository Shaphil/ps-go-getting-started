package models

type User struct {
	ID        int
	FirstName string
	LastName  string
}

var (
	users []*User
	// nextID int = 1
	nextID = 1 // the same as the above line
)
