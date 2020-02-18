package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

func main() {

	ctx, cancel := context.WithCancel(context.Background())

	fmt.Println("Chequeo de error 1:", ctx.Err())
	fmt.Println("Numero de Gorutinas 1:", runtime.NumGoroutine())

	go func() {
		n := 0

		for {
			select {
			case <-ctx.Done():
				return
			default:
				n++
				time.Sleep(time.Millisecond * 200)
				fmt.Println("Trabajando", n)
			}

		}

	}()

	time.Sleep(time.Second * 2)
	fmt.Println("Chequeo de error 2:", ctx.Err())
	fmt.Println("Numero de Gorutinas 2:", runtime.NumGoroutine())

	fmt.Println("A punto de cancelar Context.")
	cancel()
	fmt.Println("Context Cancelado.")

	time.Sleep(time.Second * 2)
	fmt.Println("Chequeo de error 3:", ctx.Err())
	fmt.Println("Numero de Gorutinas 3:", runtime.NumGoroutine())
}
