package main

import (
	"fmt"
)

func main() {
	// Arrays
	var fruitArray [2]string

	// Assign
	fruitArray[0] = "Apple"
	fruitArray[1] = "Orange"
	fmt.Println(fruitArray)

	// Declare and assign
	personArray := [2]string{"Pesho", "Kotze"}
	fmt.Println(personArray)

	// Slice
	personSlice := []string{"Pesho", "Kotze", "Pesho2"}
	fmt.Println(personSlice)
	fmt.Println(personSlice[1:2])
	fmt.Println(len(personSlice))

}
