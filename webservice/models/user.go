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

func GetUsers() []*User {
	return users
}

func AddUser(user User) (User, error) {
	user.ID = nextID
	nextID++
	users = append(users, &user)

	return user, nil
}
