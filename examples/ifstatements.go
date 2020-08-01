package examples

type User struct {
	ID        int
	FirstName string
	LastName  string
}

func ifstatements() {
	u1 := User{
		ID:        1,
		FirstName: "Stevie",
		LastName:  "Hudson",
	}
	u2 := User{
		ID:        2,
		FirstName: "Jay",
		LastName:  "Hudson",
	}
	if u1.ID == u2.ID {
		println("Same user")
	} else if u1.LastName == u2.LastName {
		println("Similar user")
	} else {
		println("Different user")
	}
}
