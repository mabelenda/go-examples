package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	defer foo()

	_, err := os.Open("sin-archivo.txt")
	if err != nil {
		log.Fatalln(err)
	}
}

func foo() {

	fmt.Println("Cuando os.Exit() es llamada, las funciones diferidas no corren")
}
