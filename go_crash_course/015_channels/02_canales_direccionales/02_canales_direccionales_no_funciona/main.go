package main

import "fmt"

func main() {

	//Channel (Canal con Buffer)
	//Channel Send only
	ca := make(chan<- int, 2)

	ca <- 42
	ca <- 43

	fmt.Println(ca)
	fmt.Println(ca)

	fmt.Printf("---------------")
	fmt.Printf("%T\n", ca)

}
