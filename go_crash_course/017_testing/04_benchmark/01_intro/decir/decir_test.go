package decir

import (
	"fmt"
	"testing"
)

func TestSaludar(t *testing.T) {

	s := Saludar("Martin")

	if s != "Bienvenido Querido Martin" {
		t.Error("Expected", "Bienvenido Querido Martin", "Got", s)
	}
}

func ExampleSaludar() {

	fmt.Println(Saludar("Martin"))

	//Output:
	//Bienvenido Querido Martin
}

func BenchmarkSaludar(b *testing.B) {

	for i := 0; i < b.N; i++ {
		Saludar("Martin")
	}
}
