package main

import "fmt"

func main() {
	type user struct {
		ID        int
		FirstName string
		LastName  string
	}

	var u user
	u.ID = 1
	u.FirstName = "Arthur"
	u.LastName = "Dent"
	fmt.Println(u)

	u2 := user{
		ID:        2,
		FirstName: "Arthur",
		LastName:  "Kent",
	}
	fmt.Println(u2)
}
