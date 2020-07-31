package examples

import "fmt"

func datatypes() {

	firstName, lastName := "Stevie", "Hudson"

	fmt.Println(firstName, lastName)

	//built in maths
	com := complex(1, 2)
	r, i := real(com), imag(com)

	fmt.Println(com)
	fmt.Println(r, i)

	//declare pointer
	var location *string = new(string)

	//dereference
	*location = "location"

	fmt.Println(location)
	fmt.Println(*location)

	//address of
	address := "address"
	pointer := &address

	fmt.Println(address)
	fmt.Println(*pointer, pointer)

	address = "new address"
	fmt.Println(*pointer, pointer)

	//constant cannot change value, must assign on same line before compile
	const pi = 3.14
	fmt.Println(pi)

	//infer datatype, evaluates each time
	const infer = 1
	fmt.Println(infer + 1)
	fmt.Println(infer + 1.2)

	//convert explict datatype
	const inferEx int = 1
	fmt.Println(float32(inferEx) + 1.2)
}
