package main

import (
	"fmt"
	"strconv"
)

// Define person struct
type Person struct {
	// firstName string
	// lastName  string
	// city      string
	// gender    string
	firstName, lastName, city, gender string
	age                               int
}

// Greeting method (value receiver)
func (person Person) greet() string {
	return "Hello, my name is " + person.firstName + " " + person.lastName + ", age: " + strconv.Itoa(person.age)
}

// HasBirthday method (pointer receiver)
func (person *Person) hasBirthday() {
	person.age++
}

// GetMarried (pointer receiver)
func (person *Person) getMarried(spouseLastName string) {
	if person.gender == "m" {
		return
	} else {
		person.lastName = spouseLastName
	}
}

func main() {
	// Init person
	person1 := Person{
		firstName: "Sam",
		lastName:  "Smith",
		city:      "NY",
		gender:    "m",
		age:       25,
	}

	person2 := Person{
		"Sanna",
		"Smith",
		"NY",
		"f",
		25,
	}

	// fmt.Println(person1, person2)
	fmt.Println(person1.lastName)
	fmt.Println(person2.age)
	person2.hasBirthday()
	person2.getMarried("Williams")
	fmt.Println(person2.greet())

}
