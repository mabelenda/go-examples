package main

import "fmt"

func main() {

	//Define map

	// emails := make(map[string]string)

	//Assing key value
	// emails["Bob"] = "Bob@gmail.com"
	// emails["Martin"] = "Martinab@gmail.com"
	// emails["Larry"] = "Larry@gmail.com"

	//declare map and add key values
	emails := map[string]string{"Bob": "Bob@gmail.com", "Martin": "Martinab@gmail.com"}

	fmt.Println(emails)
	fmt.Println(len(emails))
	fmt.Println(emails["Bob"])

	//Delete from map

	delete(emails, "Bob")
	fmt.Println(emails)
}
