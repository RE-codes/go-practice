package main

import (
	"fmt"
)

type person struct {
	name string
	fact string
	age  int
}

type secretAgent struct {
	person
	hasGun bool
}

type pracPerson struct {
	first       string
	last        string
	favIceCream []string
}

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main() {
	p1 := person{
		name: "Rich",
		fact: "Loves Go(lang)",
		age:  41,
	}

	p2 := secretAgent{
		person: person{
			name: "April",
			fact: "Loves ice cream",
			age:  39,
		},
		hasGun: true,
	}

	fmt.Println(p1, p2)

	fmt.Println(p1.name, p1.fact, p1.age)
	fmt.Println(p2.name, p2.fact, p2.age, p2.hasGun)

	// Anonymous struct w/ composite literal
	anon := struct {
		field1 string
		field2 int
		field3 bool
	}{
		field1: "Hello World",
		field2: 99,
		field3: true,
	}

	fmt.Println(anon)

	// ******** Hands on exercises ********

	person1 := pracPerson{
		first:       "April",
		last:        "Eldridge",
		favIceCream: []string{"mint chocolate chip", "cookie dough", "peanut butter cup"},
	}

	person2 := pracPerson{
		first:       "Rich",
		last:        "Eldridge",
		favIceCream: []string{"chocolate"},
	}

	fmt.Printf("person1 name: %v %v\n", person1.first, person1.last)
	fmt.Println("favorite ice creams:")
	for i, val := range person1.favIceCream {
		fmt.Printf("\t%v: %v\n", i+1, val)
	}

	fmt.Printf("person2 name: %v %v\n", person2.first, person2.last)
	fmt.Println("favorite ice creams:")
	for i, val := range person2.favIceCream {
		fmt.Printf("\t%v: %v\n", i+1, val)
	}

	// Create a map with first name as key and structs as value
	people := map[string]pracPerson{
		person1.first: person1,
		person2.first: person2,
	}

	fmt.Println(people["Rich"])
	fmt.Println(people["April"])

	fmt.Printf("person1: %v %v\n", people["April"].first, people["April"].last)
	for i, val := range people["April"].favIceCream {
		fmt.Printf("\t%v: %v\n", i+1, val)
	}

	fmt.Printf("person2: %v %v\n", people["Rich"].first, people["Rich"].last)
	for i, val := range people["Rich"].favIceCream {
		fmt.Printf("\t%v: %v\n", i+1, val)
	}

	// Better way
	for key, val := range people {
		fmt.Printf("%v %v:\n", key, val.last)
		for i, val := range val.favIceCream {
			fmt.Printf("\t%v: %v\n", i+1, val)
		}
	}

	myTruck := truck{
		vehicle: vehicle{
			doors: 2,
			color: "white",
		},
		fourWheel: true,
	}

	mySedan := sedan{
		vehicle: vehicle{
			doors: 4,
			color: "Champagne",
		},
		luxury: true,
	}

	fmt.Printf("myTruck: %v\n", myTruck)
	fmt.Printf("mySedan: %v\n", mySedan)
	fmt.Println(myTruck.color)
	fmt.Println(mySedan.luxury)

	// Create anonymous struct w/ one data type map and one data type slice
	anonymousStruct := struct {
		mapField   map[string]string
		sliceField []string
	}{
		mapField: map[string]string{
			"key1": "first value",
			"key2": "second value",
			"key3": "third value",
		},
		sliceField: []string{
			"first element",
			"second element",
			"third element",
		},
	}

	fmt.Println(anonymousStruct)
	fmt.Println(anonymousStruct.mapField)
	for key, val := range anonymousStruct.mapField {
		fmt.Printf("%v: %v\n", key, val)
	}

	for i, val := range anonymousStruct.sliceField {
		fmt.Printf("%v: %v\n", i+1, val)
	}
}
