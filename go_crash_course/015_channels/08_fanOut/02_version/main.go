package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {

	c1 := make(chan int)
	c2 := make(chan int)

	go agregar(c1)

	go fanOutIn(c1, c2)

	for v := range c2 {
		fmt.Println(v)
	}

	fmt.Println("Finalizando...")
}

func agregar(c chan<- int) {

	for i := 0; i < 100; i++ {
		c <- i
	}
	close(c)
}

func fanOutIn(c1 <-chan int, c2 chan<- int) {
	var wg sync.WaitGroup
	const GOROUTINES = 10
	wg.Add(GOROUTINES)

	for i := 0; i < GOROUTINES; i++ {
		go func() {
			for v := range c1 {
				func(v2 int) {
					c2 <- trabajoConsumeTiempo(v2)
					wg.Done()
				}(v)
			}
		}()
	}
	wg.Wait()
	close(c2)
}

func trabajoConsumeTiempo(value int) int {
	time.Sleep(time.Microsecond * time.Duration(rand.Intn(500)))
	return value + rand.Intn(1000)
}
