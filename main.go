package main

import (
	"fmt"
)

type myCustomType int // Creating custom type w/ base type int

func main() {
	var str = "hello world"
	// user _ to ignore one of the returns (either n or err, in this case)
	n, err := fmt.Println(str, 41, true) // two default returns, num "n" of bytes of value(s) and any err
	fmt.Println(n)
	fmt.Println(err)

	// Shorthand declaration operator (:=), declares and assigns, type is inferred
	x := 30
	fmt.Println(x)
	fmt.Printf("%T\n", x)  // print type of x; see docs for other formatting opts
	fmt.Printf("%#x\n", x) // prints x in hex format w/ a leading 0x (the #)

	var y myCustomType = 45
	fmt.Println(y)
	fmt.Printf("%T\n", y)

	// x = y  this doesn't work because x is of type int and y is of type myCustomType

	x = int(y) // converting value of y of type myCustomType to int and reassigning x to that value
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", y)
}
