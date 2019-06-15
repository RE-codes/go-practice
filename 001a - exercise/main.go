package main

import (
	"fmt"
	"runtime"
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

// var x int = 42
// var y string = "James Bond"
// var z bool = true

type banana int

var y int

func main() {
	// s := fmt.Sprintf("%v\t%v\t%v\t", x, y, z)
	// fmt.Println(s)
	var x banana
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)
	var temp = int(x)
	fmt.Println("temp = ", temp)
	fmt.Printf("temp is a %T\n", temp)
	y = temp
	fmt.Println("y = ", y)
	fmt.Printf("y is of type %T\n", y)
	fmt.Println(runtime.GOARCH)
	fmt.Println(runtime.GOOS)
}
