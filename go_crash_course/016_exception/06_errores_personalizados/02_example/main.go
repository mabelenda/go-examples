package main

import (
	"errors"
	"fmt"
	"log"
	"math"
)

var ErrMat = errors.New("De matematica elemental: no hay raiz cuadrada real de un numero negativo")

func main() {
	value, err := sqrt(-10)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(value)
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, ErrMat
	}
	return math.Sqrt(f), nil
}
