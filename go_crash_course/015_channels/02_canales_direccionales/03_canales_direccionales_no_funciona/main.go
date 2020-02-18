package main

import "fmt"

func main() {

	//Channel (Canal con Buffer)
	//Channel Recieve only
	ca := make(<-chan int, 2)

	fmt.Println(ca)
	fmt.Println(ca)

	fmt.Printf("---------------")
	fmt.Printf("%T\n", ca)

}
