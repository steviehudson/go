package examples

func panic() {
	//panic - when applicaiton reaches a point it cannot proceed e.g. cannot connect to the db
	//lets rest of application know it has entered a state it cannot continue to function from
	panic("Something bad just happened")

}
