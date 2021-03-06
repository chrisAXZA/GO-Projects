package main

import (
	"fmt"
)

func main() {
	// Define map
	// emails := make(map[string]string)

	// Assign
	// emails["Bob"] = "bob@gmail.com"
	// emails["Sharon"] = "sharon@gmail.com"
	// emails["Mike"] = "mike@gmail.com"

	emails := map[string]string{"Bob": "bob@gmail.com", "Sharon": "sharon@gmail.com"}

	fmt.Println(len(emails))
	fmt.Println(emails)
	fmt.Println(emails["Bob"])

	// Delete
	delete(emails, "Bob")
	fmt.Println(emails)
}
