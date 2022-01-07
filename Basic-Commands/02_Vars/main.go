package main

import "fmt"

func main() {
	// var name string = "Brad"
	name, email := "Brad", "brad@gmail.com"
	var age int32 = 37
	const isCool = true
	// isCool = false
	var size float32 = 2.3
	size2 := 1.3

	fmt.Println(name, email, age, isCool)

	fmt.Printf("%T\n", name)
	fmt.Printf("%T\n", age)
	fmt.Printf("%T\n", size)
	fmt.Printf("%T\n", size2)
}
