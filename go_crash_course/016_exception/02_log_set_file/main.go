package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	file, err := os.Create("log.txt")

	if err != nil {

		fmt.Println("Ocurrió un error", err)
	}

	defer file.Close()

	log.SetOutput(file)

	f2, err := os.Open("Sin-archivo.txt")

	if err != nil {

		log.Println("Ocurrió un error", err)
	}

	defer f2.Close()

	fmt.Println("Revisa el archivo Log.txt en el PATH")
}
