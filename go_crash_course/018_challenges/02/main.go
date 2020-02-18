package main

import (
	"fmt"

	_m "github.com/martinab/go_crash_course/018_challenges/02/math"
)

func main() {

	fmt.Println("********* SIN CHANNELS *********")
	fmt.Println("Cuadrado:", _m.Squares(20))
	fmt.Println("Cubo:", _m.Cubes(20))
	fmt.Println("Resultado sin channels:", _m.SquaresPlusCubes(20))

	fmt.Println("*********** CON CHANNELS ********")
	var numberChan = make(chan int)

	go _m.SquaresChan(20, numberChan)
	go _m.CubesChan(20, numberChan)
	go _m.SquaresPlusCubesChan(20, numberChan)

	fmt.Println("Cuadrado:", <-numberChan)

	fmt.Println("Cubo:", <-numberChan)

	fmt.Println("Resultado con channels:", <-numberChan)
}
