package main

import "fmt"

func main() {
	c := make(chan int)

	//Send
	go func() {
		for i := 0; i < 5; i++ {
			c <- i
		}
		close(c)
	}()

	//Recieve
	for v := range c {
		fmt.Println(v)
	}
	fmt.Println("Finalizando")
}
