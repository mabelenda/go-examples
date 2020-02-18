package main

import "fmt"

func main() {
	ids := []int{33, 67, 12, 312, 41, 2}

	//Loop throughs ids
	for i, id := range ids {
		fmt.Printf("%d - ID: %d\n", i, id)
	}

	//Not using index
	for _, id := range ids {
		fmt.Printf("ID: %d\n", id)
	}

	// Add ids together
	sum := 0
	for _, id := range ids {
		sum += id
	}
	fmt.Println("Sum", sum)

	//range with map
	emails := map[string]string{"Bob": "Bob@gmail.com", "Martin": "Martinab@gmail.com"}

	for k, v := range emails {
		fmt.Printf("%s: %s\n", k, v)
	}

	for k := range emails {
		fmt.Println("Name:", k)
	}
}
