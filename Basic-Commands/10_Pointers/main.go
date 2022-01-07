package main

import (
	"fmt"
)

func main() {
	a := 5
	b := &a

	fmt.Println(a, b)
	fmt.Printf("Variable -> %T\nPointer -> %T\n", a, b)

	// Use * to read val from address
	fmt.Println(b, " -> ", *b)
	fmt.Println(b, " -> ", *&a)

	// Change value with pointer
	*b = 10
	fmt.Println(a, *b)
}
