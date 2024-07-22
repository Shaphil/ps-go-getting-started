package main

type HTTPRequest struct {
	Method string
}

func main() {
	request := HTTPRequest{Method: "HEAD"}

	switch request.Method {
	case "GET":
		println("Get request")
	case "POST":
		println("Post request")
	case "DELETE":
		println("Delete request")
	case "PUT":
		println("Put request")
	default:
		println("Unhandled method")
	}
}
