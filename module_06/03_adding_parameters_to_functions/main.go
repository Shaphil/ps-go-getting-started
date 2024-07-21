package main

import "fmt"

func main() {
	port := 3000
	startWebServer(port, 2, true)
}

func startWebServer(port, numberOfRetries int, checked bool) {
	fmt.Println("starting server...")
	// do something
	fmt.Println("Server started on port:", port)
	fmt.Println("Number of retries", numberOfRetries)
	fmt.Println("Everthing checked:", checked)
}
