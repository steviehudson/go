package main

import (
	"fmt"
)

func function() {
	port := 3000

	//call function with return and arguments (_, write only variable)
	_, err := startWebServer(port)
	fmt.Println(err)
}

//function with parameters and return
func startWebServer(port int) (int, error) {
	fmt.Println("Starting server...")
	//
	fmt.Println("Server started om port", port)
	return port, nil
}
