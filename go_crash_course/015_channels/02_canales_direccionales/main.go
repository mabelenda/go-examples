package main

import "fmt"

func main() {

	//Channel (Canal con Buffer)
	cr := make(<-chan int, 2)
	cs := make(chan<- int, 2)
	ca := make(chan int, 2)

	ca <- 42
	ca <- 43

	fmt.Println(<-ca)
	fmt.Println(<-ca)
	fmt.Printf("---------------")
	fmt.Printf("%T\n", ca)

	//Especifico a general no asigna
	// ca = cs
	// ca = cr

	//Especifico a especifico no asigna
	// cs = cr

	//General a especifico si funciona
	cs = ca
	cr = ca

	fmt.Println(<-cr)

}
