package main

import "fmt"

// const (
// 	a int = 10
// 	b = 20
// )

const (
	a = 2019 + iota
	b = 2019 + iota
	c = 2019 + iota
	d = 2019 + iota
)

func main() {
	// num := 384
	// fmt.Printf("%d\t%b\t%x\n", num, num, num)

	// x := 10
	// y := 20

	// a := x == y
	// b := x <= y
	// c := x >= y
	// d := x != y
	// e := x < y
	// f := x > y

	// fmt.Println("a =", a)
	// fmt.Println("b =", b)
	// fmt.Println("c =", c)
	// fmt.Println("d =", d)
	// fmt.Println("e =", e)
	// fmt.Println("f =", f)

	// 	fmt.Println("a =", a, "b =", b)
	// 	fmt.Printf("%T\t%T\t", a, b)

	// a := 10
	// fmt.Printf("a = %d\t%b\t%x\n", a, a, a)

	// b := a << 1
	// fmt.Printf("b = %d\t%b\t%x\n", b, b, b)

	// str := `I'm a string`
	// fmt.Println(str)

	fmt.Println(a, b, c, d)
}
