package main

import "fmt"

func main() {
	c := make(chan int)

	//Send
	go send(c)

	//Recieve
	recieve(c)

	fmt.Println("Finalizando")
}

func send(c chan<- int) {
	c <- 42

}

func recieve(c <-chan int) {

	fmt.Println(<-c)
}
