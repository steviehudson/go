package controllers

import (
	"encoding/json"
	"io"
	"net/http"
)

//controllers handle routing of http request from web server to method

func RegisterControllers() {
	uc := newUserController()

	//route to user controller
	http.Handle("/users", *uc)
	http.Handle("/users/", *uc)
}

func encodeResponseAsJSON(data interface{}, w io.Writer) {
	enc := json.NewEncoder(w)
	enc.Encode(data)
}
