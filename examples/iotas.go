package examples

import "fmt"

const (
	global  = "global"
	another = "another"

	// iota const
	iotaGlobal  = iota
	iotaAnother = iota

	// math
	iotaCurrent = iota
	iotaPlus    = iota + 1
	iotaUnspecified1
	iotaBitshift = 2 << iota
	iotaUnspecified2
)

const (
	// reset iota with new const block
	iotaReset = iota
)

func iotas() {

	fmt.Println(global, another)
	fmt.Println(iotaGlobal, iotaAnother)
	fmt.Println(iotaCurrent, iotaPlus, iotaUnspecified1, iotaBitshift, iotaUnspecified2)
	fmt.Println(iotaReset)
	fmt.Println(iotaGlobal, iotaReset)
}
