package main

import "fmt"

func main() {
	par := make(chan int)
	impar := make(chan int)
	exit := make(chan int)

	//Enviar
	go enviar(par, impar, exit)

	//Recibir
	recibir(par, impar, exit)

	fmt.Println("Finalizando")

}

func enviar(p, i, e chan<- int) {
	for j := 0; j < 100; j++ {
		if j%2 == 0 {
			p <- j
		} else {
			i <- j
		}
	}

	// close(p)
	// close(i)
	e <- 0
}

func recibir(p, i, e <-chan int) {

	for {
		select {
		case v := <-p:
			fmt.Println("Desde el canal Par:", v)
		case v := <-i:
			fmt.Println("Desde el canal impar:", v)
		case v := <-e:
			fmt.Println("Desde el canal salir:", v)
			return

		}
	}
}
