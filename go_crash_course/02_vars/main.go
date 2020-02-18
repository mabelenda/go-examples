package main

import "fmt"

func main() {
	//var name = "Martin"
	var age = 21

	//Shortand
	// name := "Martin"
	name, email := "Martin", "martinab@lagash.com"

	size := 1.3

	fmt.Println(name, age, size, email)
	fmt.Printf("El tipo de dato es %T\n", size)
}
