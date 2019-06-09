package main

import (
	"fmt"
)

// func main() {
// 	x := 42
// 	y := "James Bond"
// 	z := true

// 	fmt.Println("x =", x, "y =", y, "z =", z)
// 	fmt.Println(x)
// 	fmt.Println(y)
// 	fmt.Println(z)
// }

var x int = 42
var y string = "James Bond"
var z bool = true

func main() {
	s := fmt.Sprintf("%v\t%v\t%v\t", x, y, z)
	fmt.Println(s)
}
