package main

import (
	"fmt"
	"strings"

	"github.com/martinab/go_crash_course/017_testing/04_benchmark/03_cat/mystr"
)

const s = "Nos preguntamos: ¿Quien soy yo para ser brillante, hermosa, talentosa, fabulosa?"

func main() {

	xs := strings.Split(s, " ")

	for _, v := range xs {
		fmt.Println(v)
	}

	fmt.Printf("\n%s\n", mystr.Cat(xs))
	fmt.Printf("\n%s\n", mystr.Join(xs))

}
