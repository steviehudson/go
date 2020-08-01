package examples

func forloops() {
	//forloop till condition
	var a int
	for a < 5 {
		println(a)
		a++
		if a == 3 {
			break
			//continue
		}
		//println("continuing")
	}

	//forloop till condition with post clause
	for b := 0; b < 5; b++ {
		println(b)
	}

	//infinate forloop
	var c int
	for {
		//remove if statement to run infinately
		if c == 5 {
			break
		}
		println(c)
		c++
	}

	//collection forloop
	slice := []int{1, 2, 3}
	//index and value. range keyword passes in colleciton type and returns index and value at index
	for i, v := range slice {
		println(i, v)
	}

	//loop over keys only, ignore second return value using write only variable
	wellKnownPorts := map[string]int{"http": 80, "https": 443}
	for _, v := range wellKnownPorts {
		println(v)
	}
}
