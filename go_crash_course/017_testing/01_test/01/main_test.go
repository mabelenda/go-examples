package main

import "testing"

func TestMisuma(t *testing.T) {

	v := miSuma(2, 8)
	if v != 10 {
		//t.Error("ESPERABA", VALOR ESPERADO, "OBTUVE", VALOR OBTENIDO)
		t.Error("Expected", 10, "Got", miSuma(2, 8))
	}
}
