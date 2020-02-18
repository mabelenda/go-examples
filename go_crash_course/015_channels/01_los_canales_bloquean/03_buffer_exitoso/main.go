package main

import "fmt"

func main() {

	//Channel (Canal com Buffer)
	ca := make(chan int, 1)

	ca <- 42

	fmt.Println(<-ca)
}
