package main

import "fmt"

func main() {

	//Channel (Canal con Buffer)
	ca := make(chan int, 1)

	ca <- 42
	ca <- 43

	fmt.Println(<-ca)
}
