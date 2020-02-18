package main

import (
	"fmt"

	"github.com/martinab/go_crash_course/017_testing/03_examples/01/mate"
)

func main() {

	fmt.Println(mate.Sum(2, 4, 5))
	fmt.Println(mate.Sum(4, 7, 8, 0))

}
