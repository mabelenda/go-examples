package main

import "fmt"

func main() {
	//Arrays
	// var fruitArray [2]string

	// // Assing Values
	// fruitArray[0] = "Apple"
	// fruitArray[1] = "Orange"

	//Declare and assing
	// fruitArray := [2]string{"Apple", "Orange"}

	// fmt.Println(fruitArray)

	fruitArray := []string{"Apple", "Orange", "Grape"}

	fmt.Println(fruitArray)
	fmt.Println(len(fruitArray))
	fmt.Println(fruitArray[1:2])

}
