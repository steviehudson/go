package models

type User struct {
	ID        int
	Firstname string
	LastName  string
}

var (
	users  []*User
	nextID = 1
)
