package main

import "fmt"

func main() {
	// pointer
	var firstName *string = new(string)

	// de-referencing a pointer
	*firstName = "Shaphil"
	fmt.Println(*firstName, firstName)

	lastName := "Mahmud"
	fmt.Println(lastName)

	ptr := &lastName
	fmt.Println(ptr, *ptr)

	lastName = "Ahmed"
	fmt.Println(ptr, *ptr)
}
