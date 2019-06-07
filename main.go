package main

import (
	"fmt"
)

func main() {
	var str = "hello world"
	// user _ to ignore one of the returns (either n or err, in this case)
	n, err := fmt.Println(str, 41, true)
	fmt.Println(n)
	fmt.Println(err)
}
