package main

import (
	"net/http"

	"shaphil.me/webservice/controllers"
)

func main() {
	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)
}
