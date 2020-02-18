package main

import (
	"os"
)

func main() {

	_, err := os.Open("sin-archivo.txt")
	if err != nil {
		panic(err)
	}
}
