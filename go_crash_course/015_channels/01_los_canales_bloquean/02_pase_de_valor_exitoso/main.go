package main

import "fmt"

func main() {

	// Unbuffered Channel (Canal sin Buffer)
	ca := make(chan int)

	go func() {
		ca <- 42
	}()

	fmt.Println(<-ca)
}
