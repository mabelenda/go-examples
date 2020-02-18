package mystr

import (
	"fmt"
	"strings"
	"testing"
)

func TestJoin(t *testing.T) {
	s := "Shaken not stirred"
	xs := strings.Split(s, " ")

	s = Join(xs)

	if s != "Shaken not stirred" {
		t.Error("Got", s, "Want", "Shaken not stirred")
	}

}

func TestCat(t *testing.T) {
	s := "Shaken not stirred"
	xs := strings.Split(s, " ")

	s = Cat(xs)

	if s != "Shaken not stirred" {
		t.Error("Got", s, "Want", "Shaken not stirred")
	}

}

func ExampleCat() {

	s := "Shaken not stirred"
	xs := strings.Split(s, " ")

	fmt.Println(Cat(xs))

	//Output:
	//Shaken not stirred
}

func ExampleJoin() {

	s := "Shaken not stirred"
	xs := strings.Split(s, " ")

	fmt.Println(Join(xs))

	//Output:
	//Shaken not stirred
}

const s = "Nos preguntamos: ¿Quien soy yo para ser brillante, hermosa, talentosa, fabulosa?"

func BenchmarkCat(b *testing.B) {
	xs := strings.Split(s, " ")
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		Cat(xs)
	}
}

func BenchmarkJoin(b *testing.B) {
	xs := strings.Split(s, " ")
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		Cat(xs)
	}
}
