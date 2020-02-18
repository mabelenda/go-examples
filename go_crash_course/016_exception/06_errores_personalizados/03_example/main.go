package main

import (
	"fmt"
	"log"
	"math"
)

type ubicacionError struct {
	lat  string
	long string
	err  error
}

func (n ubicacionError) Error() string {
	return fmt.Sprintf("Un error de concepto matematico ha ocurrido en: %v %v %v", n.lat, n.long, n.err)
}

func main() {
	_, err := sqrt(-10.23)
	if err != nil {
		log.Println(err)
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		nme := fmt.Errorf("Error Mat: raíz cuadrada de número negativo: %v", f)
		return 0, ubicacionError{"50.2289 N", "994556 W", nme}
	}
	return math.Sqrt(f), nil
}
