package main

import (
	"fmt"
	"net/http"

	"shaphil.me/webservice/controllers"
)

func main() {
	controllers.RegisterControllers()

	host := "localhost"
	port := "3000"
	hostname := host + ":" + port
	fmt.Println("Started server at:", hostname)

	http.ListenAndServe(":3000", nil)
}
