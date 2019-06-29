package main

import (
	"fmt"
)

func main() {
	// for i := 0; i < 10; i++ {
	// 	fmt.Printf("The outer loop: %v\n", i)
	// 	for j := 0; j < 10; j++ {
	// 		fmt.Printf("\tThe inner loop: %v\n", j)
	// 	}
	// }

	// x := "something" // scoped at the func main level
	// if x == "something" {
	// 	fmt.Println("yup!")
	// }

	if x := "something"; x == "something else" { // scoped to just the if statement
		fmt.Println("yup!")
	} else {
		fmt.Println("nope!")
	}
}
