package main

import "fmt"

func main() {
	par := make(chan int)
	impar := make(chan int)
	exit := make(chan bool)

	//Enviar
	go enviar(par, impar, exit)

	//Recibir
	recibir(par, impar, exit)

	fmt.Println("Finalizando")

}

func enviar(p, i chan<- int, e chan<- bool) {
	for j := 0; j < 100; j++ {
		if j%2 == 0 {
			p <- j
		} else {
			i <- j
		}
	}
	// close(p)
	// close(i)
	close(e)
}

func recibir(p, i <-chan int, e <-chan bool) {

	for {
		select {
		case v := <-p:
			fmt.Println("Desde el canal Par:", v)
		case v := <-i:
			fmt.Println("Desde el canal impar:", v)
		//Comma Ok (Se obtiene el valor del channel y un boolean que identifica si el channel tiene datos o no [Si esta cerrado])
		case v, ok := <-e:

			if !ok {
				fmt.Println("Desde comma Ok, se deberia obtener 0 porque el channel no tiene valores, se obtiene:", v)
				return
			} else {
				fmt.Println("Desde comma Ok, se deberia obtener un valor porque el channel tiene valores, se obtiene:", v)

			}
		}
	}
}
