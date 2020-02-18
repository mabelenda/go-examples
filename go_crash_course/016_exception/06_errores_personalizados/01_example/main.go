package main

import (
	"errors"
	"fmt"
	"log"
	"math"
)

func main() {

	// fmt.Println("Seleccione un nuevo para aplicarle su raiz cuadrada:")

	// var number float64
	// fmt.Scanf("%d", number)

	value, err := sqrt(-10)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(value)
}

func sqrt(f float64) (float64, error) {

	if f < 0 {
		return 0, errors.New("De matematica elemental: no hay raiz cuadrada real de un numero negativo")
	}
	return math.Sqrt(f), nil
}
