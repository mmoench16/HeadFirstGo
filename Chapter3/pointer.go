package main

import (
	"fmt"
)

func negate(myBoolean *bool) {
	*myBoolean = !*myBoolean
}

func main() {
	var myInt int
	var myIntPointer *int

	myInt = 42
	myIntPointer = &myInt
	fmt.Println(*myIntPointer)

	truth := true
	negate(&truth)
	fmt.Println(truth)
}
