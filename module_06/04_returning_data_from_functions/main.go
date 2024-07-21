package main

import (
	"errors"
	"fmt"
)

func main() {
	port := 3000
	_, err := startWebServer(port)
	fmt.Println(err)
}

func startWebServer(port int) (int, error) {
	fmt.Println("starting server...")
	// do something
	fmt.Println("Server started on port:", port)

	return port, errors.New("Error: Something went horribly wrong")
}
