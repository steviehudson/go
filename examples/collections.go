package examples

import "fmt"

func collections() {

	//arrays - fixed
	var arr [3]int
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3

	fmt.Println(arr)
	fmt.Println(arr[0])

	arr2 := [3]int{1, 2, 3}

	fmt.Println(arr2)

	//slices - dynamic
	//create from array
	slice := arr[:]
	fmt.Println(slice)

	//slice like pointer and will update array
	arr[1] = 42
	slice[2] = 27
	fmt.Println(arr, slice)

	//go to manage underlying array
	s1 := []int{1, 2, 3}
	fmt.Println(s1)

	//slice is not fixed size
	s1 = append(s1, 4)
	fmt.Println(s1)

	//create slice from slice
	// add index :
	s2 := slice[1:]
	s3 := slice[:2]
	s4 := slice[1:2]
	fmt.Println(s2, s3, s4)

	//maps - key value
	m := map[string]int{"foo": 42}
	fmt.Println(m)
	fmt.Println(m["foo"])

	m["foo"] = 27
	fmt.Println(m)

	delete(m, "foo")
	fmt.Println(m)

	//structs  - limited alternative to class (data side)
	type user struct {
		ID        int
		FirstName string
		LastName  string
	}
	var u user
	fmt.Println(u)

	u.ID = 1
	u.FirstName = "Stevie"
	u.LastName = "Hudson"
	fmt.Println(u)
	fmt.Println(u.FirstName)

	u2 := user{ID: 1,
		FirstName: "Stephanie",
		LastName:  "Hudson",
	}
	fmt.Println(u2)
}
