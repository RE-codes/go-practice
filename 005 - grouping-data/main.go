package main

import (
	"fmt"
)

func main() {
	mySlice := []int{2, 4, 6, 8, 10} // so this is how you make a slice using composite literal
	fmt.Printf("The slice: %v\n", mySlice)

	// regular for loop using len function
	for i := 0; i < len(mySlice); i++ {
		fmt.Printf("The value at index %v is %v\n", i, mySlice[i])
	}

	// using for...range to do the same thing
	for i, val := range mySlice {
		fmt.Printf("The value at position %v is %v\n", i, val)
	}

	// slicing a slice
	sliceOfmySlice := mySlice[1:3] // should include values from position 1 up to position 3 i.e. [4, 6]
	fmt.Printf("The slice of the slice is: %v\n", sliceOfmySlice)

	// appending to a slice
	myAppendedSlice := append(mySlice, 12, 14, 16)
	fmt.Printf("After appending: %v\n", myAppendedSlice)
	// *or*
	toAppend := []int{12, 14, 16}
	myAppendedSlice = append(mySlice, toAppend...)
	fmt.Printf("Appending a slice to a slice: %v\n", myAppendedSlice)

	// stick it in the middle by slicing to the index and appending the slice from that point onward
	toInsert := toAppend
	firstPart := mySlice[:3] // beginning UP TO index 3
	lastPart := mySlice[3:]  // index 3 to the end of the slice
	firstPart = append(firstPart, toInsert...)
	sliceWithValuesInserted := append(firstPart, lastPart...)
	fmt.Printf("Trying to stick it in the middle: %v\n", sliceWithValuesInserted)
	// docs in github wiki indicate a better way to do this using "copy"

	// delete from a slice
	oldSlice := append(sliceWithValuesInserted[:3], sliceWithValuesInserted[6:]...)
	fmt.Printf("Deleting the inserted values: %v\n", oldSlice)

	// multi-dimensional slice aka slice of slices
	s1 := []string{"one", "two", "three"}
	s2 := []string{"four", "five", "six"}

	mds := [][]string{s1, s2}
	fmt.Printf("A multi-dimensional slice: %v\n", mds)

	// MAPS
	myMap := map[string]int{
		"Rich":  41,
		"April": 39, // ***Include trailing comma***
	}
	fmt.Printf("My map: %v\n", myMap)
	fmt.Printf("If I want the value associated with the key, Rich: %v\n", myMap["Rich"])
	fmt.Printf("If I want the value associated with the key, April: %v\n", myMap["April"])
	fmt.Printf("If I look up a key that is not in the map, like \"Harrison\": %v\n", myMap["Harrison"])
	// Defaults to "zero" value if key does not exist

	// Comma-OK syntax
	val, ok := myMap["Harrison"]
	fmt.Printf("Using comma-ok method, val = %v and ok = %v\n", val, ok)

	// To check if key exists:
	testVal := "Harrison"
	if val, ok := myMap[testVal]; ok {
		fmt.Println("This isn't going to print", val) // because ok will evaluate to "false"
	} else {
		fmt.Printf("%v does not exist\n", testVal)
	}

	testVal = "April"
	if val, ok := myMap[testVal]; ok {
		fmt.Printf("This does print the value for %v: %v\n", testVal, val)
	} else {
		fmt.Printf("%v does not exist\n", testVal)
	}

	// To add a key:value to a map
	myMap["Harrison"] = 6

	fmt.Printf("When I add Harrison: %v\n", myMap)

	// Looping through with range
	for key, val := range myMap {
		fmt.Printf("The value for key, %v, is %v\n", key, val)
	}

	// Delete an entry from a map
	delete(myMap, "Harrison")

	for key, val := range myMap {
		fmt.Printf("The value for key, %v, is %v\n", key, val)
	}

	delete(myMap, "Gwen") // Will not throw error if item does not exist in map

	// ****** Hands on exercises ******

	fmt.Println("****** Hands on starts here ******")

	myArr := [5]int{1, 2, 3, 4, 5}

	for _, val := range myArr {
		fmt.Printf("%v\n", val)
	}

	fmt.Printf("%T\n\n", myArr)

	mySliceExercise := []int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}

	for _, val := range mySliceExercise {
		fmt.Printf("%v\n", val)
	}

	fmt.Printf("%T\n\n", mySliceExercise)

	slc1 := mySliceExercise[:5]  // Expect [2, 4, 6, 8, 10]
	slc2 := mySliceExercise[5:]  // Expect [12, 14, 16, 18, 20]
	slc3 := mySliceExercise[2:7] // Expect [6, 8, 10, 12, 14]
	slc4 := mySliceExercise[1:6] // Expect [4, 6, 8, 10, 12]

	fmt.Printf("slc1 = %v\n", slc1)
	fmt.Printf("slc2 = %v\n", slc2)
	fmt.Printf("slc3 = %v\n", slc3)
	fmt.Printf("slc4 = %v\n\n", slc4)

	mySliceExercise = append(mySliceExercise, 52)
	fmt.Println(mySliceExercise)

	mySliceExercise = append(mySliceExercise, 53, 54, 55)
	fmt.Println(mySliceExercise)

	y := []int{56, 57, 58, 59, 60}
	mySliceExercise = append(mySliceExercise, y...)
	fmt.Println(mySliceExercise)

	mySliceExercise = []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	mySliceExercise = append(mySliceExercise[:3], mySliceExercise[6:]...)
	fmt.Println(mySliceExercise)

	states := []string{
		"Alabama",
		"Alaska",
		"Arizona",
		"Arkansas",
		"California",
		"Colorado",
		"Connecticut",
		"Delaware",
		"Florida",
		"Georgia",
		"Hawaii",
		"Idaho",
		"Illinois",
		"Indiana",
		"Iowa",
		"Kansas",
		"Kentucky",
		"Louisiana",
		"Maine",
		"Maryland",
		"Massachusetts",
		"Michigan",
		"Minnesota",
		"Mississippi",
		"Missouri",
		"Montana",
		"Nebraska",
		"Nevada",
		"New Hampshire",
		"New Jersey",
		"New Mexico",
		"New York",
		"North Carolina",
		"North Dakota",
		"Ohio",
		"Oklahoma",
		"Oregon",
		"Pennsylvania",
		"Rhode Island",
		"South Carolina",
		"South Dakota",
		"Tennessee",
		"Texas",
		"Utah",
		"Vermont",
		"Virginia",
		"Washington",
		"West Virginia",
		"Wisconsin",
		"Wyoming",
	}

	fmt.Println(len(states))
	fmt.Println(cap(states))

	for i := 0; i < len(states); i++ {
		fmt.Printf("%v\t%v\n", i, states[i])
	}

	bond := []string{"James", "Bond", "Shaken, not stirred"}
	moneypenny := []string{"Miss", "Moneypenny", "Hellooooo, James"}

	people := [][]string{bond, moneypenny}

	for _, val := range people {
		fmt.Println(val)
	}

	for i := range people {
		for _, val := range people[i] {
			fmt.Println(val)
		}
	}

	pracMap := map[string][]string{
		"bond_james": []string{"Shaken, not stirred", "Martinis", "Women"},
		"moneypenny": []string{"James Bond", "Literature", "Computer Science"},
		"no_dr":      []string{"Being evil", "Ice cream", "Sunsets"},
	}

	for key, val := range pracMap {
		fmt.Printf("key: %v\t\tval: %v\n", key, val)
		for i, val := range pracMap[key] {
			fmt.Printf("\tindex: %v\t value: %v\n", i, val)
		}
	}

	pracMap["evil_dr"] = []string{"Frickin laser beams", "kittteeeesss", "one biiiillllllion dollars"}

	for key, val := range pracMap {
		fmt.Printf("key: %v\t favs: %v\n", key, val)
	}

	delete(pracMap, "no_dr")

	fmt.Println("**************************")
	for key, val := range pracMap {
		fmt.Printf("key: %v\t favs: %v\n", key, val)
	}
}
