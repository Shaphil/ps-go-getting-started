package models

import (
	"errors"
	"fmt"
)

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
	if user.ID != 0 {
		return User{}, errors.New("new user must not include id or it must be set to zero")
	}

	user.ID = nextID
	nextID++
	users = append(users, &user)

	return user, nil
}

func GetUserById(id int) (User, error) {
	for _, user := range users {
		if user.ID == id {
			return *user, nil
		}
	}

	return User{}, fmt.Errorf("User with ID `%v` not found", id)
}

func UpdateUser(user User) (User, error) {
	for i, candidate := range users {
		if candidate.ID == user.ID {
			users[i] = &user
			return user, nil
		}
	}

	return User{}, fmt.Errorf("User with ID `%v` not found", user.ID)
}

func RemoveUserById(id int) error {
	for i, user := range users {
		if user.ID == id {
			users = append(users[:i], users[i+1:]...)
			return nil
		}
	}

	return fmt.Errorf("User with ID `%v` not found", id)
}
