package examples

type HTTPRequest struct {
	Method string
}

func switch() {
	r := HTTPRequest{Method: "GET"}

	switch r.Method {
	case "GET":
		println("GET Request")
		//falls into method below
		fallthrough
	case "DELETE":
		println("DELETE Request")
	case "POST":
		println("POST Request")
	case "PUT":
		println("PUT Request")
	default:
		println("Unhandled method ")
	}
}