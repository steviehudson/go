package main

import (
	"net/http"

	"github.com.pluralsight/webservice/controllers"
)

//fire up webservice

func main() {
	//front controller - handle all routing in entire project using calls to handle methds in RegisterControllers method
	controllers.RegisterControllers()
	//http server up and running
	http.ListenAndServe(":3000", nil)
}
