package main

import (
	"fmt"
	"strconv"
)

//Define person struct
type Person struct {
	// firstName string
	// lastName  string
	// city      string
	// gender    string
	// age       int

	firstName, lastName, city, gender string
	age                               int
}

// Greeting method (value reciver)
func (p Person) greet() string {
	return "Hello, my name is " + p.firstName + " " + p.lastName + " and I am " + strconv.Itoa(p.age)
}

//HasBirthday method (pointer reciver)
func (p *Person) hasBirthday() {
	p.age++
}

//GetMarried (pointer reciver)
func (p *Person) getMarried(spouseLastName string) {

	if p.gender == "Female" {
		p.lastName = spouseLastName
	}
}

func main() {

	// Init person using struct
	person1 := Person{
		firstName: "Martin",
		lastName:  "Abelenda",
		city:      "Lomas",
		gender:    "Male",
		age:       21,
	}

	// Alternative
	person2 := Person{
		"Laura",
		"Ricardes",
		"Lomas",
		"Female",
		23,
	}

	// fmt.Println(person1.firstName)

	// person1.age++

	fmt.Println(person1)

	person1.hasBirthday()

	person1.getMarried("Rodriguez")

	person2.getMarried("Perez")

	fmt.Println(person1.greet())

	fmt.Println(person2.greet())

}
