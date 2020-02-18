package main

import "fmt"

func main() {
	//Long method
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i++
	}

	//Short method
	for i := i; i <= 10; i++ {
		fmt.Printf("Numer %d\n", i)
	}

	//FizzzBuzz
	for i := 1; i <= 100; i++ {
		if i%15 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)

		}
	}
}
