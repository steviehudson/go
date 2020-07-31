package main

import (
	"fmt"

	"github.com.pluralsight/webservice/models"
)

func main() {
	u := models.User{
		ID:        1,
		Firstname: "Stevie",
		LastName:  "Hudson",
	}
	fmt.Println(u)
}
